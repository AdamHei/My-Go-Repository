package exchangeApis

import (
	"log"
	"encoding/json"
	"net/http"
)

const NUMEXCHANGES = 4

func FetchAllExchanges(ch chan<- map[string]map[string]interface{}) {
	go fetchBidAskPoloniex(ch)
	go fetchBidAskGemini(ch)
	go fetchBidAskKraken(ch)
	go fetchBidAskGDAX(ch)
}

func fetchBidAskGemini(ch chan<- map[string]map[string]interface{}) {
	resp, err := http.Get(GEMINIURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	geminiResponse := new(GeminiTicker)
	err = json.NewDecoder(resp.Body).Decode(geminiResponse)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp.Body.Close()

	ch <- geminiResponse.GetExchangeData()
}

func fetchBidAskKraken(ch chan<- map[string]map[string]interface{}) {
	resp, err := http.Get(KRAKENURL)
	if err != nil {
		log.Fatal("Kraken request failed", err.Error())
	}

	krakenResponse := new(KrakenResponse)
	err = json.NewDecoder(resp.Body).Decode(krakenResponse)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp.Body.Close()

	ch <- krakenResponse.GetExchangeData()
}

func fetchBidAskGDAX(ch chan<- map[string]map[string]interface{}) {
	resp, err := http.Get(GDAXURL)
	if err != nil {
		log.Fatal("GDAX request failed", err.Error())
	}

	gdaxResponse := new(GDAXTicker)
	err = json.NewDecoder(resp.Body).Decode(gdaxResponse)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp.Body.Close()

	ch <- gdaxResponse.GetExchangeData()
}

func fetchBidAskPoloniex(ch chan<- map[string]map[string]interface{}) {
	resp, err := http.Get(POLONIEXURL)
	if err != nil {
		log.Fatal("Poloniex request failed", err.Error())
	}

	fullResponse := new(map[string]PoloniexTicker)
	err = json.NewDecoder(resp.Body).Decode(fullResponse)

	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	btcTicker := (*fullResponse)["USDT_BTC"]
	ch <- btcTicker.GetExchangeData()
}