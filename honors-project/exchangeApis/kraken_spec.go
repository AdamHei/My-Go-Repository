package exchangeApis

import (
	"net/http"
	"log"
	"encoding/json"
)

const krakenurl = "https://api.kraken.com/0/public/Ticker?pair=XBTUSD"

func fetchBidAskKraken(ch chan<- map[string]map[string]string) {
	resp, err := http.Get(krakenurl)
	if err != nil {
		log.Println("Could not fetch Kraken data: ", err)
	}

	krakenResponse := new(KrakenResponse)
	err = json.NewDecoder(resp.Body).Decode(krakenResponse)

	resp.Body.Close()
	if err != nil {
		log.Println("Could not parse Kraken json: ", err)
		return
	}

	ch <- krakenResponse.GetExchangeData()
}

func (response KrakenResponse) GetExchangeData() map[string]map[string]string {
	return map[string]map[string]string{
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
