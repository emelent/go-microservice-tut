BIN_DIR=build
SERVER_BIN=${BIN_DIR}/accountservice
HEALTH_BIN=${BIN_DIR}/healthcheck
export CGO_ENABLED=0

all: cmd/server/main.go cmd/healthchecker/main.go
	go build -o ${SERVER_BIN} cmd/server/main.go
	go build -o ${HEALTH_BIN} cmd/healthchecker/main.go

docker-images:
	docker build -t emelent/accountservice .
	
docker-deploy:
	docker service rm accountservice
	docker service create --name=accountservice --replicas=1\
		--network=my_network -p=6767:6767 emelent/accountservice

docker-clean-deploy:
	docker service create --name=accountservice --replicas=1\
		--network=my_network -p=6767:6767 emelent/accountservice

clean:
	rm ${BIN_DIR}/*

run-server:
	go run cmd/server/main.go

run-health:
	go run cmd/healthchecker/main.go -port=6767

test-all:
	go test github.com/emelent/go-microservice-tut/accounts/pkg/...
test-service:
	go test github.com/emelent/go-microservice-tut/accounts/pkg/service
