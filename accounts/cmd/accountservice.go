package main

import (
	"fmt"

	"github.com/emelent/go-microservice-tut/accounts/pkg/dbclient"
	"github.com/emelent/go-microservice-tut/accounts/pkg/service"
)

const appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenDb()
	service.DBClient.Seed()
}
