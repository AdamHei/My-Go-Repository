package exchangeApis

import "time"

const GDAXURL = "https://api.gdax.com/products/BTC-USD/ticker"

func (response GDAXTicker) GetBidAsk() []string {
	return []string{"GDAX", response.Bid, response.Ask}
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
