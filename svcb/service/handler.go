package service

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/iyabchen/go-zipkin-example/model"
)

// GetQuotes sleeps 3 seconds and return a fixed quote object.
func GetQuotes(w http.ResponseWriter, r *http.Request) {

	time.Sleep(2 * time.Second)

	var accountID = mux.Vars(r)["accountID"]

	quote := model.Quote{
		Language: "en",
		ServedBy: "svcb",
		Text:     "May the source be with account ID " + accountID,
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(quote)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
