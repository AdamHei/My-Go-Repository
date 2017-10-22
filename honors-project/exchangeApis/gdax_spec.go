package exchangeApis

import (
	"time"
	"net/http"
	"log"
	"encoding/json"
)

const gdaxurl = "https://api.gdax.com/products/BTC-USD/ticker"

func fetchBidAskGDAX(ch chan<- map[string]map[string]string) {
	resp, err := http.Get(gdaxurl)
	if err != nil {
		log.Println("Could not fetch data from GDAX: ", err)
		return
	}

	gdaxResponse := new(GDAXTicker)
	err = json.NewDecoder(resp.Body).Decode(gdaxResponse)

	resp.Body.Close()
	if err != nil {
		log.Println("Could not parse GDAX json: ", err)
		return
	}

	ch <- gdaxResponse.GetExchangeData()
}

func (response GDAXTicker) GetExchangeData() map[string]map[string]string {
	return map[string]map[string]string{
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
