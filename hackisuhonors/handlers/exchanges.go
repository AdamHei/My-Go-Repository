package handlers

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/adamhei/hackisuhonors/exchangeApis"
)

func AllBidAskData(writer http.ResponseWriter, _ *http.Request) {
	ch := make(chan map[string]map[string]interface{})

	exchangeApis.FetchAllExchanges(ch)
	response := make(map[string]map[string]interface{})

	for i := 0; i < exchangeApis.NUMEXCHANGES; i++ {
		exchangeData := <-ch
		for key, value := range exchangeData {
			response[key] = value
		}
	}
	respond(writer, response, nil)
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
