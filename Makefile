
all:
	docker-compose build

re-deploy-accounts:
	docker service rm accountservice
	docker service create --name=accountservice --replicas=1\
		--network=my_network -p=6767:6767 emelent/accountservice

deploy-accounts:
	docker service create --name=accountservice --replicas=1\
		--network=my_network -p=6767:6767 emelent/accountservice