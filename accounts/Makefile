TARGET_OS=linux
TARGET_EXE=accountservice-linux-amd64
export GOOS=${TARGET_OS}

docker-build:
	go build -o ${TARGET_EXE} cmd/main.go
	export GOOS=darwin

