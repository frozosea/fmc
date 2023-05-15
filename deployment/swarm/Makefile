del-network:
	docker network rm public
network:
	docker network create -d overlay --attachable public
dashboard:
	curl -L https://downloads.portainer.io/ce2-14/portainer-agent-stack.yml -o portainer-agent-stack.yml && docker stack deploy -c portainer-agent-stack.yml portainer
user-api:
	docker stack deploy -c services/user-api/user-api.yaml user
tracking-api:
	docker stack deploy -c services/tracking/tracking.yaml tracking
gateway:
	docker stack deploy -c services/gateway/gateway.yaml gateway
schedule-tracking:
	docker stack deploy -c services/schedule-tracking/schedule-tracking.yaml schedule_tracking
freight-service:
	docker stack deploy -c services/freight-service/freight-service.yaml freight_service
del-user:
	docker stack rm user
del-tracking:
	docker stack rm tracking
del-schedule-tracking:
	docker stack rm schedule_tracking
del-portainer:
	docker stack rm portainer
del-gateway:
	docker stack rm gateway
del-all: del-user del-tracking del-schedule-tracking del-portainer del-gateway
deploy: network dashboard tracking-api user-api schedule-tracking gateway
