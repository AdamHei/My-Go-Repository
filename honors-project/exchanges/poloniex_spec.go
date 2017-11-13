package exchanges

import (
	"net/http"
	"encoding/json"
)

const poloniexurl = "http://poloniex.com/public?command=returnTicker"

func fetchBidAskPoloniex(ch chan<- LimitedJson) {
	resp, err := http.Get(poloniexurl)
	if err != nil {
		ErrorHandler("Could not fetch Poloniex data: "+err.Error(), ch)
		return
	}

	fullResponse := new(map[string]PoloniexTicker)
	err = json.NewDecoder(resp.Body).Decode(fullResponse)
	resp.Body.Close()

	if err != nil {
		ErrorHandler("Could not parse Poloniex json: "+err.Error(), ch)
		return
	}

	btcTicker := (*fullResponse)["USDT_BTC"]
	ch <- btcTicker.GetExchangeData()
}

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
