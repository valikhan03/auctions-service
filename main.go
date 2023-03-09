package main

import (
	"google.golang.org/grpc"
	"net"
	"log"
	"fmt"

	pb_mng_srvc "github.com/valikhan03/auctions-service/pb/auctions_management_service"
	pb_run_srvc "github.com/valikhan03/auctions-service/pb/auctions_run_service"
	"github.com/valikhan03/auctions-service/elasticsearch"
	"github.com/valikhan03/auctions-service/redis"
	"github.com/valikhan03/auctions-service/services"
	"github.com/valikhan03/auctions-service/repositories"
	"github.com/valikhan03/auctions-service/models"
)

func main() {
	models.InitConfigs()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", models.ServerConfigGlobal.Host, models.ServerConfigGlobal.Port))
	if err != nil{
		log.Fatalf("Listener error: %s", err.Error())
	}

	server := grpc.NewServer()
	elasticConn := elasticsearch.InitElasticsearchConn()
	redisConn := redis.InitRedisConn()

	mng_repo := repositories.NewManageRepository(redisConn, elasticConn)
	run_repo := repositories.NewRunRepository(redisConn)

	mng_srvc := services.NewAuctionsManagementService(mng_repo)
	run_srvc := services.NewAuctionsRunService(run_repo)

	pb_mng_srvc.RegisterAuctionsManagementServiceServer(server, mng_srvc)
	pb_run_srvc.RegisterAuctionsRunServiceServer(server, run_srvc)


	log.Println("Starting auctions-service...")
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
