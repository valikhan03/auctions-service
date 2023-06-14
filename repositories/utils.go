package repositories

import (
	"context"

	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/valikhan03/tool"
)

func (r *ManageRepository) GetAuctionType(auctionID string) (int, error) {
	req := esapi.GetRequest{
		Index:      tool.AuctionsIDX,
		DocumentID: auctionID,
		Pretty:     true,
	}

	response, err := req.Do(context.TODO(), r.elastic)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	var res map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		log.Printf("Service.GetAuction: %x\n", err)
		return -1, err
	}

	return res["_source"].(map[string]interface{})["type"].(int), nil
}

func (r *ManageRepository) GetAllLotsList(auctionID string) ([]string, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":    fmt.Sprintf("auction_id=%s", auctionID),
				"fields":   []string{"title", "description"},
				"type":     "best_fields",
				"operator": "and",
			},
		},
	}

	var buffer bytes.Buffer

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		log.Fatalf("Service.Search: %s", err.Error())
	}

	res, err := r.elastic.Search(
		r.elastic.Search.WithIndex(tool.LotsIDX),
		r.elastic.Search.WithBody(&buffer),
		r.elastic.Search.WithTrackTotalHits(true),
		r.elastic.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Service.Search: %s", err.Error())
	}

	resbody := make(map[string]interface{})
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&resbody)
	if err != nil {
		log.Fatalf("Service.Search: %s", err.Error())
	}

	if res.IsError() {
		log.Printf("[%s] %s: %s",
			res.Status(),
			resbody["error"].(map[string]interface{})["type"],
			resbody["error"].(map[string]interface{})["reason"],
		)
		return nil, fmt.Errorf("[%s] %s: %s", res.Status(), resbody["error"].(map[string]interface{})["type"], resbody["error"].(map[string]interface{})["reason"])
	}

	fmt.Println(resbody)
	var lots []string

	for _, hit := range resbody["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]
		lots = append(lots, source.(map[string]interface{})["id"].(string))
	}

	return lots, nil
}

func CreateInvoice() {

}
