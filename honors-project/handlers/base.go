package handlers

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/adamhei/honors-project/exchangeApis"
)

func Index(writer http.ResponseWriter, _ *http.Request) {
	respond(writer, "Welcome to the Bitcoin Arbitrage Detector", nil)
}

func respond(writer http.ResponseWriter, data interface{}, err *exchangeApis.MyError) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		log.Println(err.Error())
		http.Error(writer, err.Error(), err.ErrorCode())
	} else {
		json.NewEncoder(writer).Encode(data)
	}
}
