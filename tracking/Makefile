migrate-up:
	migrate -path ./schema -database 'postgres://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE_NAME}?sslmode=disable' up
migrate-down:
	migrate -path ./schema -database 'postgres://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE_NAME}?sslmode=disable' down
test-unit:
	go test -parallel 15 -v --short ./...
run:
	go run cmd/main.go