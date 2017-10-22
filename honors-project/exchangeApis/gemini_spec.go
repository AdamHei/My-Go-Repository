package exchangeApis

import (
	"net/http"
	"log"
	"encoding/json"
)

const geminiurl = "https://api.gemini.com/v1/pubticker/btcusd"

func fetchBidAskGemini(ch chan<- map[string]map[string]string) {
	resp, err := http.Get(geminiurl)
	if err != nil {
		log.Println("Could not fetch Gemini data: ", err)
	}

	geminiResponse := new(GeminiTicker)
	err = json.NewDecoder(resp.Body).Decode(geminiResponse)

	resp.Body.Close()
	if err != nil {
		log.Println("Could not parse Gemini json: ", err)
		return
	}

	ch <- geminiResponse.GetExchangeData()
}

func (response GeminiTicker) GetExchangeData() map[string]map[string]string {
	return map[string]map[string]string{
		"Gemini": {
			"Bid": response.Bid,
			"Ask": response.Ask,
		},
	}
}

type GeminiTicker struct {
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	BTCVolume volume `json:"volume"`
	Last      string `json:"last"`
}

type volume struct {
	BTC       string `json:"BTC"`
	USD       string `json:"USD"`
	Timestamp int    `json:"timestamp"`
}
