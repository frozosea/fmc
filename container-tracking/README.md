# Container tracking microservice

Steps to run this project:

1. create .env file and set this variables:
- REQUEST_TIMEOUT_MS
- WAIT_SELECTOR_TIMEOUT
- POSTGRES_HOST
- POSTGRES_PORT
- POSTGRES_USER
- POSTGRES_PASSWORD
- POSTGRES_DATABASE
- REDIS_URL
- CONTAINER_TRACKING_RESULT_REDIS_TTL_SECONDS
- GRPC_PORT
3. Setup database settings inside `data-source.ts` file
4. Run `npm start` command
