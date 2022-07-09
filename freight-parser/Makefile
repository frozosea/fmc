include .env

migrate-up:
	migrate -path ./schema -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DATABASE}?sslmode=disable' up
migrate-down:
	migrate -path ./schema -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DATABASE}?sslmode=disable' down
run:
	go run cmd/main.go
proto:
	protoc --proto_path=${PWD}/pkg/proto/  --go_out=${PWD}/pkg/proto/  --go-grpc_out=${PWD}/pkg/proto/   -I ${PWD}/pkg/proto/ freight.proto