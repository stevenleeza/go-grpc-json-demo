protobuf:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative homeaffairspb/homeaffairs.proto

run-server:
	go run server/server.go

run-client:
	go run client/client.go

run: 
	docker-compose up -d --build

stop: 
	docker-compose down