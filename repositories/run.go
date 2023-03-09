package repositories

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/redis/go-redis/v9"
)

type RunRepository struct {
	mx  *sync.RWMutex
	rdb *redis.Client
}

func NewRunRepository(redisConn *redis.Client) *RunRepository {
	return &RunRepository{
		mx:  &sync.RWMutex{},
		rdb: redisConn,
	}
}

func (r *RunRepository) SetSuggestPrice(ctx context.Context, auctionID, lotID string, userID, price int64) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	key := fmt.Sprintf("%s/%s", auctionID, lotID)
	val := map[string]interface{}{"userid": userID, "price": price}
	err := r.rdb.HSet(ctx, key, val).Err()
	if err != nil {
		return err
	}

	err = r.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: auctionID,
		Values: map[string]interface{}{
			"user_id": userID,
			"lot_id":  lotID,
			"price":   price,
		},
	}).Err()

	if err != nil {
		return err
	}

	//r.rdb.SAdd()

	err = r.rdb.LPush(ctx, fmt.Sprintf("%s/%s", auctionID, lotID),
		map[string]interface{}{"userid": userID, "price": price}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RunRepository) GetMaxPrice(ctx context.Context, auctionID, lotID string) (int64, error) {

	res, err := r.rdb.HGet(ctx, fmt.Sprintf("%s/%s", auctionID, lotID), "price").Result()
	if err != nil {
		return -1, err
	}
	price, err := strconv.Atoi(res)
	if err != nil{
		return -1, err
	}
	return int64(price), nil
}

func (r *RunRepository) ReadMaxPrice(ctx context.Context, auctionID, lotID, userID string) (map[string]string, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	xstr, err := r.rdb.XReadGroup(context.TODO(), &redis.XReadGroupArgs{
		Group:    lotID,
		Consumer: userID,
		Streams:  []string{auctionID, ">"},
	}).Result()

	if err != nil {
		return nil, err
	}

	res := make(map[string]string)

	for _, m := range xstr[0].Messages {
		if m.Values["lot_id"].(string) == lotID {
			res["user_id"] = m.Values["user_id"].(string)
			res["price"] = m.Values["price"].(string)
			r.rdb.XAck(context.TODO(), auctionID, lotID, m.ID)
			break
		}
	}

	return res, nil
}

func (r *RunRepository) InitConsumer(userID, auctionID, lotID string) error {
	return r.rdb.XGroupCreateConsumer(context.TODO(), auctionID, lotID, userID).Err()
}
