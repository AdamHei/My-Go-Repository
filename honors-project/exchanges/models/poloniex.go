package models

func (response PoloniexTicker) GetExchangeData() LimitedJson {
	return LimitedJson{
		"Poloniex": {
			"Bid": response.Bid,
			"Ask": response.Ask,
		},
	}
}

type PoloniexTicker struct {
	Id            int    `json:"id"`
	LastPrice     string `json:"last"`
	Ask           string `json:"lowestAsk"`
	Bid           string `json:"highestBid"`
	PercentChange string `json:"percentChange"`
	BaseVolume    string `json:"baseVolume"`
	QuoteVolume   string `json:"quoteVolume"`
	IsFrozen      string `json:"isFrozen"`
	High24Hr      string `json:"high24hr"`
	Low24Hr       string `json:"low24hr"`
}
