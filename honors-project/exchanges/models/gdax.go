package models

import (
	"time"
)

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
