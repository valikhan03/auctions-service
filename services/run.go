package services

import (
	"context"
	"log"

	pb "github.com/valikhan03/auctions-service/pb/auctions_run_service"

	"github.com/valikhan03/auctions-service/repositories"
)

type AuctionsRunService struct {
	repository *repositories.RunRepository
}


func NewAuctionsRunService(run_repo *repositories.RunRepository) *AuctionsRunService {
	return &AuctionsRunService{
		repository: run_repo,
	}
}

func (s *AuctionsRunService) SuggestPrice(ctx context.Context, req *pb.SuggestPriceRequest) (*pb.SuggestPriceResponse, error) {
	maxprice, err := s.repository.GetMaxPrice(ctx, req.AuctionID, req.LotID)
	if err != nil{
		log.Printf("SuggestPrice error: %s",err.Error())
		return nil, err
	}

	if req.NewPrice <= maxprice{
		return &pb.SuggestPriceResponse{Status: 0, Error: "Suggestion is lower than current"}, nil
	}

	err = s.repository.SetSuggestPrice(ctx, req.AuctionID, req.LotID, req.UserID, req.NewPrice)
	if err != nil {
		log.Printf("SuggestPrice error: %s",err.Error())
		return nil, err
	}

	return &pb.SuggestPriceResponse{Status: 200}, nil
}

func (s *AuctionsRunService) GetCurrentPrice(req *pb.GetCurrentPriceRequest, 
					stream pb.AuctionsRunService_GetCurrentPriceServer) error {
	err := s.repository.InitConsumer(req.UserID, req.AuctionID, req.LotID)
	if err != nil{
		log.Printf("InitConsumer error: %s", err.Error())
		return err
	}

	for{
		res, err := s.repository.ReadMaxPrice(context.TODO(), req.AuctionID, req.LotID, req.UserID)
		if err != nil{
			log.Printf("GetCurrentPrice error: %s",err.Error())
			break
		}
		stream.Send(&pb.GetCurrentPriceResponse{UserPrice: res})
	}

	return err
}

