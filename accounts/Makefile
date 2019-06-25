TARGET_EXE=accountservice
export CGO_ENABLED=0

build:
	go build -o ${TARGET_EXE} cmd/accountservice.go

test-all:
	go test github.com/emelent/go-microservice-tut/accounts/pkg/...
test-service:
	go test github.com/emelent/go-microservice-tut/accounts/pkg/service