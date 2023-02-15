gen_proto: 
	protoc -I . --go_out=. --go-grpc_out=. protobuf/auctions_management_service.proto
	protoc -I . --go_out=. --go-grpc_out=. protobuf/auctions_run_service.proto

