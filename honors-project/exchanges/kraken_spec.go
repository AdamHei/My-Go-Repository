package exchanges

import (
	"net/http"
	"encoding/json"
)

const krakenurl = "https://api.kraken.com/0/public/Ticker?pair=XBTUSD"

func fetchBidAskKraken(ch chan<- LimitedJson) {
	resp, err := http.Get(krakenurl)
	if err != nil {
		ErrorHandler("Could not fetch Kraken data:"+err.Error(), ch)
	}

	krakenResponse := new(KrakenResponse)
	err = json.NewDecoder(resp.Body).Decode(krakenResponse)

	resp.Body.Close()
	if err != nil {
		ErrorHandler("Could not parse Kraken json:"+err.Error(), ch)
		return
	}

	ch <- krakenResponse.GetExchangeData()
}

func (response KrakenResponse) GetExchangeData() LimitedJson {
	return LimitedJson{
		"Kraken": {
			"Bid": response.Result.XBTUSD.BidArr[0],
			"Ask": response.Result.XBTUSD.AskArr[0],
		},
	}
}

type KrakenResponse struct {
	Error  []string `json:"error"`
	Result Result   `json:"result"`
}

type Result struct {
	XBTUSD PairBody `json:"XXBTZUSD"`
}

type PairBody struct {
	AskArr               []string `json:"a"`
	BidArr               []string `json:"b"`
	ClosedTradeArr       []string `json:"c"`
	VolumeArr            []string `json:"v"`
	VolumeWeightedAvgArr []string `json:"p"`
	NumTradesArr         []int    `json:"t"`
	LowArr               []string `json:"l"`
	HighArr              []string `json:"h"`
	OpenPrice            string   `json:"o"`
}
