deploy: network dashboard deploy-user-api deploy-tracking-api deploy-gateway
network:
	docker network create -d overlay --attachable public
dashboard:
	curl -L https://downloads.portainer.io/ce2-14/portainer-agent-stack.yml -o portainer-agent-stack.yml && docker stack deploy -c portainer-agent-stack.yml portainer
deploy-user-api:
	docker stack deploy -c services/user-api/user-api.yaml user
deploy-tracking-api:
	docker stack deploy -c services/tracking/tracking.yaml tracking
deploy-gateway:
	docker stack deploy -c services/gateway/gateway.yaml gateway
