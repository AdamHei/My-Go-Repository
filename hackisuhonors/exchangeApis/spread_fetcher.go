package exchangeApis

import (
	"log"
	"encoding/json"
	"net/http"
)

const NUMEXCHANGES = 4

func FetchAllExchanges(ch chan<- map[string]map[string]string) {
	go fetchBidAskPoloniex(ch)
	go fetchBidAskGemini(ch)
	go fetchBidAskKraken(ch)
	go fetchBidAskGDAX(ch)
}

func fetchBidAskGemini(ch chan<- map[string]map[string]string) {
	resp, err := http.Get(GEMINIURL)
	if err != nil {
		log.Println(err)
	}

	geminiResponse := new(GeminiTicker)
	err = json.NewDecoder(resp.Body).Decode(geminiResponse)

	resp.Body.Close()
	if err != nil {
		log.Println(err)
	}

	ch <- geminiResponse.GetExchangeData()
}

func fetchBidAskKraken(ch chan<- map[string]map[string]string) {
	resp, err := http.Get(KRAKENURL)
	if err != nil {
		log.Println(err)
	}

	krakenResponse := new(KrakenResponse)
	err = json.NewDecoder(resp.Body).Decode(krakenResponse)

	resp.Body.Close()
	if err != nil {
		log.Println(err)
	}

	ch <- krakenResponse.GetExchangeData()
}

func fetchBidAskGDAX(ch chan<- map[string]map[string]string) {
	resp, err := http.Get(GDAXURL)
	if err != nil {
		log.Println(err)
		return
	}

	gdaxResponse := new(GDAXTicker)
	err = json.NewDecoder(resp.Body).Decode(gdaxResponse)

	resp.Body.Close()
	if err != nil {
		log.Println(err)
	}

	ch <- gdaxResponse.GetExchangeData()
}

func fetchBidAskPoloniex(ch chan<- map[string]map[string]string) {
	resp, err := http.Get(POLONIEXURL)
	if err != nil {
		log.Println(err)
		return
	}

	fullResponse := new(map[string]PoloniexTicker)
	err = json.NewDecoder(resp.Body).Decode(fullResponse)

	resp.Body.Close()
	if err != nil {
		log.Println(err)
	}

	btcTicker := (*fullResponse)["USDT_BTC"]
	ch <- btcTicker.GetExchangeData()
}
