package exchanges

import (
	"time"
	"net/http"
	"encoding/json"
)

const gdaxurl = "https://api.gdax.com/products/BTC-USD/ticker"

func fetchBidAskGDAX(ch chan<- LimitedJson) {
	resp, err := http.Get(gdaxurl)
	if err != nil {
		ErrorHandler("Could not fetch data from GDAX:"+err.Error(), ch)
		return
	}

	gdaxResponse := new(GDAXTicker)
	err = json.NewDecoder(resp.Body).Decode(gdaxResponse)
	resp.Body.Close()

	if err != nil {
		ErrorHandler("Could not parse GDAX json:"+err.Error(), ch)
		return
	}

	ch <- gdaxResponse.GetExchangeData()
}

func (response GDAXTicker) GetExchangeData() LimitedJson {
	return LimitedJson{
		"GDAX": {
			"Bid": response.Bid,
			"Ask": response.Ask,
		},
	}
}

type GDAXTicker struct {
	LastTradeId    int       `json:"trade_id"`
	LastTradePrice string    `json:"price"`
	LastTradeSize  string    `json:"size"`
	Bid            string    `json:"bid"`
	Ask            string    `json:"ask"`
	Volume         string    `json:"volume"`
	Time           time.Time `json:"time"`
}
