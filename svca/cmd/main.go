package main

import (
	"fmt"

	"github.com/iyabchen/go-zipkin-example/svca/data"
	"github.com/iyabchen/go-zipkin-example/svca/service"
)

var appName = "account-service"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initDB()
	service.StartWebServer("8080")
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initDB() {
	service.DBClient = &data.BoltClient{}
	service.DBClient.OpenDb()
	service.DBClient.Seed()
}
