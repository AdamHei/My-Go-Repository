package exchanges

import (
	"net/http"
	"encoding/json"
	"fmt"
	"time"
)

const geminiTickerUrl = "https://api.gemini.com/v1/pubticker/btcusd"
const geminiHistoryUrl = "https://api.gemini.com/v1/trades/btcusd?since=%d&limit_trades=500"

func fetchBidAskGemini(ch chan<- LimitedJson) {
	resp, err := http.Get(geminiTickerUrl)
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

func GetTradeHistory(from time.Time, _ time.Time) []Order {
	formattedUrl := fmt.Sprintf(geminiHistoryUrl, from.Unix())
	resp, err := http.Get(formattedUrl)
	if err != nil {
		fmt.Println(err)
		return make([]Order, 0)
	}

	orders := make([]Order, 0)
	err = json.NewDecoder(resp.Body).Decode(&orders)
	resp.Body.Close()

	if err != nil {
		fmt.Println(err)
		return make([]Order, 0)
	}

	return orders
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

type Order struct {
	Timestamp   int    `json:"timestamp"`
	TimestampMs int    `json:"timestampms"`
	TID         int    `json:"tid"`
	Price       string `json:"price"`
	Amount      string `json:"amount"`
	Exchange    string `json:"exchange"`
	Type        string `json:"type"`
}
