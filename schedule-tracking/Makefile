proto:
	protoc --proto_path=${PWD}/pkg/proto/  --go_out=${PWD}/pkg/proto/  --go-grpc_out=${PWD}/pkg/proto/   -I ${PWD}/pkg/proto/ schedule_tracking.proto
run:
	go run cmd/main.go