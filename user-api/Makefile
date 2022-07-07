include .env

migrate-up:
	migrate -path ./schema -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@0.0.0.0:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' up
migrate-down:
	migrate -path ./schema -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@0.0.0.0:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' down
test-unit:
	go test -v --short ./...
proto:
	protoc --proto_path=${PWD}/pkg/proto/  --go_out=${PWD}/pkg/proto/  --go-grpc_out=${PWD}/pkg/proto/   -I ${PWD}/pkg/proto/ user.proto
run:
	go run cmd/main.go