package main

import (
	"fmt"

	"github.com/iyabchen/go-zipkin-example/svcb/service"
)

var appName = "quote-service"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("8080")
}
