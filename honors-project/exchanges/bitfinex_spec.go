package exchanges

import (
	"net/http"
	"encoding/json"
)

const bitfinexurl = "https://api.bitfinex.com/v1/pubticker/btcusd"

func fetchBidAskBitfinex(ch chan<- LimitedJson) {
	resp, err := http.Get(bitfinexurl)
	if err != nil {
		ErrorHandler("Could not fetch data from Bitfinex:"+err.Error(), ch)
		return
	}

	bitResponse := new(BitfinexTicker)
	err = json.NewDecoder(resp.Body).Decode(bitResponse)

	resp.Body.Close()
	if err != nil {
		ErrorHandler("Could not parse Bitfinex response"+err.Error(), ch)
		return
	}

	ch <- bitResponse.GetExchangeData()
}

func (response BitfinexTicker) GetExchangeData() LimitedJson {
	return LimitedJson{
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
