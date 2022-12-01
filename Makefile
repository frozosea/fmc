user-proto:
	protoc --proto_path=${PWD}/protobuf/user  --go_out=${PWD}/protobuf/user   --go-grpc_out=${PWD}/protobuf/user    -I ${PWD}/protobuf/user user.proto

schedule-tracking-proto:
	protoc --proto_path=${PWD}/protobuf/schedule-tracking   --go_out=${PWD}/protobuf/schedule-tracking    --go-grpc_out=${PWD}/protobuf/schedule-tracking    -I ${PWD}/protobuf/schedule-tracking  schedule_tracking.proto

freight-proto:
	protoc --proto_path=${PWD}/protobuf/freight   --go_out=${PWD}/protobuf/freight    --go-grpc_out=${PWD}/protobuf/freight    -I ${PWD}/protobuf/freight  freight.proto

update-proto:
	cd schedule-tracking;\
	go mod tidy;\
	go get github.com/frozosea/fmc-pb@latest;\
	cd ..;\
	cd freight-service;\
	go mod tidy;\
	go get github.com/frozosea/fmc-pb@latest;\
	cd ..;\
    cd user-api;\
    go mod tidy;\
  	go get github.com/frozosea/fmc-pb@latest;\
    cd ..;\
    cd api-gateway;\
    go mod tidy;\
  	go get github.com/frozosea/fmc-pb@latest;\

