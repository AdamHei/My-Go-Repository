package exchanges

import (
	"net/http"
	"encoding/json"
)

const geminiurl = "https://api.gemini.com/v1/pubticker/btcusd"

func fetchBidAskGemini(ch chan<- LimitedJson) {
	resp, err := http.Get(geminiurl)
	if err != nil {
		ErrorHandler("Could not fetch Gemini data:"+err.Error(), ch)
	}

	geminiResponse := new(GeminiTicker)
	err = json.NewDecoder(resp.Body).Decode(geminiResponse)

	resp.Body.Close()
	if err != nil {
		ErrorHandler("Could not parse Gemini json:"+err.Error(), ch)
		return
	}

	ch <- geminiResponse.GetExchangeData()
}

func (response GeminiTicker) GetExchangeData() LimitedJson {
	return LimitedJson{
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
