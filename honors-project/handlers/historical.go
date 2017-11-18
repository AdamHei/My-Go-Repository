package handlers

import (
	"github.com/adamhei/honors-project/exchanges/gemini"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// Given exchange1 and exchange2, return a historical comparison
func Compare(writer http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)
	_ = args["exchange1"]
	_ = args["exchange2"]

	orders := gemini.GetTradeHistory(time.Date(2017, time.January, 1, 12, 0, 0, 0, time.Local), time.Now())

	respond(writer, orders, nil)
}
