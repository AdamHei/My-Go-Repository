package exchangeApis

import (
	"net/http"
	"log"
	"encoding/json"
)

const bitfinexurl = "https://api.bitfinex.com/v1/pubticker/btcusd"

func fetchBidAskBitfinex(ch chan<-map[string]map[string]string)  {
	resp, err := http.Get(bitfinexurl)
	if err != nil {
		log.Println("Could not fetch data from Bitfinex: ", err)
		return
	}

	bitResponse := new(BitfinexTicker)
	err = json.NewDecoder(resp.Body).Decode(bitResponse)

	resp.Body.Close()
	if err != nil {
		log.Println("Could not parse Bitfinex response: ", err)
		return
	}

	ch <- bitResponse.GetExchangeData()
}

func (response BitfinexTicker) GetExchangeData() map[string]map[string]string {
	return map[string]map[string]string{
		"Bitfinex": {
			"Bid": response.Bid,
			"Ask": response.Ask,
		},
	}
}

type BitfinexTicker struct {
	Mid       string `json:"mid"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastPrice string `json:"last_price"`
	Low       string `json:"low"`
	High      string `json:"high"`
	Volume    string `json:"volume"`
	Timestamp string `json:"timestamp"`
}
