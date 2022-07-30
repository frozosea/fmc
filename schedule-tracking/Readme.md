## Schedule tracking microservice

### Env variables

- TRACKING_GRPC_HOST
- TRACKING_GRPC_PORT
- USER_GRPC_HOST
- USER_GRPC_PORT
- SENDER_NAME
- SENDER_EMAIL
- UNISENDER_API_KEY
- POSTGRES_USER
- POSTGRES_PASSWORD
- POSTGRES_DATABASE
- POSTGRES_HOST
- POSTGRES_PORT

### How to run:

    make run 

### Go to `conf/cfg.ini` and change parameters if you want.

Methods are in `pkg/proto`.

Default timezone of running tasks is Asia/Vladivostok. If you want to change it you can open dockerfile and change the timezone.