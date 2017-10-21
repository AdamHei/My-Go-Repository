package exchangeApis

import "time"

const GDAXURL = "https://api.gdax.com/products/BTC-USD/ticker"

func (response GDAXTicker) GetExchangeData() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
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
