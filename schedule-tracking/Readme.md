## Schedule tracking microservice

Do not scale or redeploy this application,because if you do it, all schedule tasks will be deleted,and you should parse
logs and in hand mode add it.

### Env variables

- TRACKING_GRPC_HOST
- TRACKING_GRPC_PORT
- USER_GRPC_HOST
- USER_GRPC_PORT
- SENDER_NAME
- SENDER_EMAIL
- UNISENDER_API_KEY

### How to run:

    make run 

### Go to `conf/cfg.ini` and change parameters if you want.


Methods are in `pkg/proto`.

### TODO

- write script for parse logs and automatically reset deleted tasks.