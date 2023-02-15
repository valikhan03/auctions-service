package services

import (
	"bytes"
	"context"
	"encoding/json"

	//"github.com/valikhan03/tool"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/redis/go-redis/v9"
	pb "github.com/valikhan03/auctions-service/pb/auctions_management_service"
)

type AuctionsManagementService struct {
	rdb     *redis.Client
	elastic *elasticsearch.Client
}

func InitAuctionsManagementService() {

}

func (s *AuctionsManagementService) StartAuction(ctx context.Context, req *pb.StartAuctionRequest) (*pb.StartAuctionResponse, error) {
	//set status ACTIVE
	//add products and start prices to Redis
	//allow user sessions to achieve data - auth level
	status, err := json.Marshal(map[string]interface{}{"STATUS":"ACTIVE"})

	ereq := esapi.UpdateRequest{
		Index: "auctions",
		DocumentID: req.AuctionId,
		Body: bytes.NewBuffer(status),
	}

	res, err := ereq.Do(ctx, s.elastic)
	if err != nil{

	}

	if res.IsError(){

	}

	defer res.Body.Close()

	dataReq := esapi.GetRequest{
		Index: "auctions",
		DocumentID: req.AuctionId,
	}
	
	res, err = dataReq.Do(ctx, s.elastic)
	if err != nil{

	}

	//add to redis


	return &pb.StartAuctionResponse{Status: 200}, nil
}

func (s *AuctionsManagementService) CancelAuction(ctx context.Context, req *pb.CancelAuctionRequest) (*pb.CancelAuctionResponse, error) {
	//set status CANCELLED
	//migrate data from Redis to ELK and Postgres

	return &pb.CancelAuctionResponse{Status: 200}, nil
}

func (s *AuctionsManagementService) SetStartTime(ctx context.Context, req *pb.SetStartTimeRequest) (*pb.SetStartTimeResponse, error) {
	return nil, nil
}
