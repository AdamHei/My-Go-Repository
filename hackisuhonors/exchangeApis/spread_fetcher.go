package exchangeApis

import (
	"log"
	"encoding/json"
	"net/http"
)

const NUMEXCHANGES = 4

func FetchAllExchanges(ch chan<- []string) {
	go fetchBidAskPoloniex(ch)
	go fetchBidAskGemini(ch)
	go fetchBidAskKraken(ch)
	go fetchBidAskGDAX(ch)
}

func fetchBidAskGemini(ch chan<- []string) {
	resp, err := http.Get(GEMINIURL)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	geminiResponse := new(GeminiTicker)
	err = json.NewDecoder(resp.Body).Decode(geminiResponse)
	if err != nil {
		log.Fatal(err.Error())
	}

	ch <- geminiResponse.GetBidAsk()
}

func fetchBidAskKraken(ch chan<- []string) {
	resp, err := http.Get(KRAKENURL)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("Kraken request failed", err.Error())
	}

	krakenResponse := new(KrakenResponse)
	err = json.NewDecoder(resp.Body).Decode(krakenResponse)
	if err != nil {
		log.Fatal(err.Error())
	}

	ch <- krakenResponse.GetBidAsk()
}

func fetchBidAskGDAX(ch chan<- []string) {
	resp, err := http.Get(GDAXURL)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("GDAX request failed", err.Error())
	}

	gdaxResponse := new(GDAXTicker)
	err = json.NewDecoder(resp.Body).Decode(gdaxResponse)
	if err != nil {
		log.Fatal(err.Error())
	}

	ch <- gdaxResponse.GetBidAsk()
}

func fetchBidAskPoloniex(ch chan<- []string) {
	resp, err := http.Get(POLONIEXURL)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("Poloniex request failed", err.Error())
	}

	fullResponse := new(map[string]PoloniexTicker)
	err = json.NewDecoder(resp.Body).Decode(fullResponse)

	if err != nil {
		log.Fatal(err)
	}

	btcTicker := (*fullResponse)["USDT_BTC"]

	ch <- btcTicker.GetBidAsk()
}