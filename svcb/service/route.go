package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iyabchen/go-zipkin-example/svcb/common"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

// Route defines a single route, e.g. a human readable name, HTTP method and the
// pattern the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the type Routes which is just an array (slice) of Route struts.
type Routes []Route

// Initialize our routes
var routes = Routes{

	Route{
		"GetQuotes",
		"GET",
		"/quotes/{accountID}",
		GetQuotes,
	},
}

// NewRouter returns a pointer to a mux.Router we can use as a handler.
func NewRouter() *mux.Router {

	// Create an instance of the Gorilla router
	router := mux.NewRouter().StrictSlash(true)
	router.Use(zipkinhttp.NewServerMiddleware(
		common.GetTracer(),
		zipkinhttp.SpanName("svcb-mw")), // name for request span
	)
	// Iterate over the routes we declared in routes.go and attach them to the router instance
	for _, route := range routes {

		// Attach each route, uses a Builder-like pattern to set each route up.
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
