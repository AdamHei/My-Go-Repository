package models

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
