swag:
	swag init -g cmd/main.go
tracking-proto:
	protoc --proto_path=../container-tracking/src/server/proto/server.proto  --go_out=${PWD}/tracking/ --go-grpc_out=${PWD}/tracking/   -I ${PWD}/container-tracking/src/server/proto/ user.proto
run:
	go run cmd/main.go