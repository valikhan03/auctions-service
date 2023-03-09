package services

import (
	"context"
	"log"

	pb "github.com/valikhan03/auctions-service/pb/auctions_management_service"
	"github.com/valikhan03/auctions-service/repositories"
)

type AuctionsManagementService struct {
	repository *repositories.ManageRepository
}

func NewAuctionsManagementService(mng_repo *repositories.ManageRepository) *AuctionsManagementService {
	return &AuctionsManagementService{
		repository: mng_repo,
	}
}

func (s *AuctionsManagementService) StartAuction(ctx context.Context, req *pb.StartAuctionRequest) (*pb.StartAuctionResponse, error) {
	//set status ACTIVE
	//add products and start prices to Redis
	//allow user sessions to achieve data - auth level
	//create stream in redis
	err := s.repository.SetStatusActive(req.AuctionId)
	if err != nil{
		log.Printf("SetStatusActive error: %s", err.Error())
		return nil, err
	}

	err = s.repository.MigrateLotsRedis(req.AuctionId)
	if err != nil{
		log.Printf("MigrateLotsRedis error: %s", err.Error())
		return nil, err
	}

	return &pb.StartAuctionResponse{Status: 200}, nil
}

func (s *AuctionsManagementService) CancelAuction(ctx context.Context, req *pb.CancelAuctionRequest) (*pb.CancelAuctionResponse, error) {
	//set status CANCELLED
	//migrate data from Redis to Postgres
	//send data to Payment Service
	err := s.repository.SetStatusCancelled(req.AuctionId)
	if err != nil{
		log.Printf("SetStatusCancelled error: %s", err.Error())
		return nil, err
	}

	return &pb.CancelAuctionResponse{Status: 200}, nil
}

func (s *AuctionsManagementService) SetStartTime(ctx context.Context, req *pb.SetStartTimeRequest) (*pb.SetStartTimeResponse, error) {
	return nil, nil
}
