package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/redis/go-redis/v9"
	"github.com/valikhan03/auctions-service/kafka"
	"github.com/valikhan03/tool"
)

type CmdMessage struct {
}

type ManageRepository struct {
	rdb       *redis.Client
	elastic   *elasticsearch.Client
	eventChan chan kafka.Event
}

func NewManageRepository(redisConn *redis.Client, elasticConn *elasticsearch.Client) *ManageRepository {
	return &ManageRepository{
		rdb:     redisConn,
		elastic: elasticConn,
	}
}

func (r *ManageRepository) SetStatusActive(auctionID string) error {
	//set status ACTIVE
	//add products and start-prices to Redis
	//allow user sessions to achieve data - auth level
	//create stream in redis

	body := map[string]interface{}{
		"doc": map[string]interface{}{
			"status": tool.ACTIVE,
		},
	}
	status, err := json.Marshal(body)
	if err != nil {
		return err
	}

	updReq := esapi.UpdateRequest{
		Index:      tool.AuctionsIDX,
		DocumentID: auctionID,
		Body:       bytes.NewReader(status),
	}

	updRes, err := updReq.Do(context.TODO(), r.elastic)
	if err != nil {
		return err
	}

	if updRes.IsError() {
		return fmt.Errorf("[%s] %s", updRes.Status(), updRes.String())
	}

	return nil
}

func (r *ManageRepository) MigrateLotsRedis(auctionID string) error {
	//add to redis
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"auction_id": auctionID,
			},
		},
	}

	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return err
	}

	searchRes, err := r.elastic.Search(
		r.elastic.Search.WithContext(context.TODO()),
		r.elastic.Search.WithIndex(tool.LotsIDX),
		r.elastic.Search.WithBody(&buffer),
		r.elastic.Search.WithTrackTotalHits(true),
		r.elastic.Search.WithPretty(),
	)
	if err != nil {
		return err
	}
	defer searchRes.Body.Close()

	resbody := map[string]interface{}{}
	err = json.NewDecoder(searchRes.Body).Decode(&resbody)
	if err != nil {
		return err
	}
	docs := resbody["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, doc := range docs {
		source := doc.(map[string]interface{})["_source"].(map[string]interface{})
		lotID := fmt.Sprintf("%d", int(source["id"].(float64)))
		key := fmt.Sprintf("%s/%s", source["auction_id"], lotID)
		val := map[string]interface{}{"userid": "-", "price": source["start_price"]}
		err := r.rdb.HSet(context.TODO(), key, val).Err()
		if err != nil {
			return err
		}

		err = r.rdb.XAdd(context.TODO(), &redis.XAddArgs{
			Stream: auctionID,
			Values: map[string]interface{}{
				"user_id": "-",
				"lot_id":  lotID,
				"price":   source["start_price"],
			},
		}).Err()
		if err != nil {
			return err
		}

		err = r.rdb.XGroupCreate(context.TODO(), auctionID, lotID, "0").Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ManageRepository) SetStatusCancelled(auctionID string) error {
	//set status CANCELLED
	//migrate data from Redis to Postgres
	//send data to Payment Service

	status, err := json.Marshal(map[string]interface{}{"status": tool.CANCELLED})
	if err != nil {
		return err
	}

	updReq := esapi.UpdateRequest{
		Index:      tool.AuctionsIDX,
		DocumentID: auctionID,
		Body:       bytes.NewBuffer(status),
	}

	updRes, err := updReq.Do(context.TODO(), r.elastic)
	if err != nil {
		return err
	}

	if updRes.IsError() {
		return errors.New(fmt.Sprintf("[%s] %s", updRes.Status(), updRes.String()))
	}

	lots, err := r.GetAllLotsList(auctionID)
	if err != nil {
		return err
	}

	if len(lots) == 0 {
		return errors.New("No Lots found")
	}

	for _, lotID := range lots {
		price, err := r.rdb.HGet(context.TODO(), fmt.Sprintf("%s/%s", auctionID, lotID), "price").Result()
		if err != nil {
			return err
		}
		currency, err := r.rdb.HGet(context.TODO(), fmt.Sprintf("%s/%s", auctionID, lotID), "currency").Result()
		if err != nil {
			return err
		}
		// price, err := strconv.Atoi(price_res)
		// if err != nil{
		// 	return err
		// }

		userid, err := r.rdb.HGet(context.TODO(), fmt.Sprintf("%s/%s", auctionID, lotID), "userid").Result()
		if err != nil {
			return err
		}
		// userid, err := strconv.Atoi(userid_res)
		// if err != nil{
		// 	return err
		// }

		r.eventChan <- kafka.Event{
			Command: tool.CRE_INV_EVENT,
			Entity:  map[string]interface{}{"price": price, 
											"currency": currency,
											"user_id": userid, 
											"product_id": lotID,
											"product_type":3, 
											"auction_id": auctionID},
		}
	}

	return nil
}


// func (r *ManageRepository) CheckRequirements(auction_id string) (bool, error) {
// 	query := `select p_is_payed from tb_invoices
// 			  where exists(select 1 from tb_invoices t where
// 						   t.product_type=$1 and
// 						   t.product_id=$2)`
	
	
// }