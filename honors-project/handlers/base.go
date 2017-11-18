package handlers

import (
	"encoding/json"
	"github.com/adamhei/honors-project/exchanges/errorhandling"
	"log"
	"net/http"
)

func Index(writer http.ResponseWriter, _ *http.Request) {
	respond(writer, "{Welcome to the Bitcoin Arbitrage Detector}", nil)
}

func respond(writer http.ResponseWriter, data interface{}, err *errorhandling.MyError) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		log.Println(err.Error())
		http.Error(writer, err.Error(), err.Code)
	} else {
		json.NewEncoder(writer).Encode(data)
	}
}
