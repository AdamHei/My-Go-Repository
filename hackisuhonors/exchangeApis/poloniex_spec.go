package exchangeApis

const POLONIEXURL = "http://poloniex.com/public?command=returnTicker"

func (response PoloniexTicker) GetExchangeData() map[string]map[string]string {
	return map[string]map[string]string{
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
