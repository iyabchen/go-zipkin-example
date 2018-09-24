package service

import (
	"log"
	"net/http"
)

// StartWebServer starts the default http web server on the given port
func StartWebServer(port string) {

	r := NewRouter()
	// set default http default handler
	http.Handle("/", r)

	// start default http server
	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Println("An error occurred starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
