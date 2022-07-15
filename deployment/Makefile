ui-dashboard:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.6.0/aio/deploy/recommended.yaml && kubectl apply -f dashboard/
proxy:
	kubectl proxy
token:
	 kubectl -n kubernetes-dashboard create token admin-user
postgres-instance:
	kubectl apply -f postgres/
redis-instance:
	kubectl apply -f redis/
gateway:
	kubectl apply -f services/gateway/
user-app:
	kubectl apply -f services/user-api/
tracking:
	kubectl apply -f services/tracking/
services: tracking user-app gateway
