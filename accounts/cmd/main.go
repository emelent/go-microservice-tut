package main

import (
	"fmt"

	"github.com/emelent/go-microservice-tut/accounts/pkg/service"
)

const appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}
