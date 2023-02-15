package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	redis "github.com/redis/go-redis/v9"
	pb "github.com/valikhan03/auctions-service/pb/auctions_run_service"
)

type AuctionsRunService struct {
	db redisClient
}

var maxprice int64

type redisClient struct{
	MX sync.RWMutex
	rdb *redis.Client
	consumerGroup *redis.XReadGroupArgs
}

func InitAuctionsRunService() {

}

func (s *AuctionsRunService) SuggestPrice(ctx context.Context, req *pb.SuggestPriceRequest) (*pb.SuggestPriceResponse, error) {
	if req.NewPrice <= maxprice{
		return &pb.SuggestPriceResponse{Status: 0, Error: "Suggestion is lower than current"}, nil
	}

	s.db.MX.Lock()
	cmd := s.db.rdb.HSet(ctx, fmt.Sprintf("%s/%s", req.AuctionID, req.ProductID), 
		map[string]interface{}{"userid": req.UserID, "price": req.NewPrice})
	s.db.MX.Unlock()

	if cmd.Err() != nil{
	}

	maxprice = req.NewPrice

	err := s.db.rdb.XAdd(ctx, &redis.XAddArgs{
		Values: map[string]interface{}{

		},
	}).Err()
	
	if err != nil{
		
	}

	return &pb.SuggestPriceResponse{Status: 200}, nil
}

func (s *AuctionsRunService) GetCurrentPrice(req *pb.GetCurrentPriceRequest, stream pb.AuctionsRunService_GetCurrentPriceServer) error {
	for{
		
		str, err := s.db.rdb.XReadGroup(context.TODO(), &redis.XReadGroupArgs{
			Group: "",
			Consumer: "",
			Streams: []string{},
			Count: 1,
			Block: time.Second,
		}).Result()
		
		if err != nil{

		}

		for _, m := range str[0].Messages {
			stream.Send(&pb.GetCurrentPriceResponse{Username: m.Values["username"].(string), Price: m.Values["price"].(string)})
		}

	}
}