package exchangeApis

import (
	"net/http"
	"log"
	"encoding/json"
)

const poloniexurl = "http://poloniex.com/public?command=returnTicker"

func fetchBidAskPoloniex(ch chan<- map[string]map[string]string) {
	resp, err := http.Get(poloniexurl)
	if err != nil {
		log.Println("Could not fetch Poloniex data: ", err)
		return
	}

	fullResponse := new(map[string]PoloniexTicker)
	err = json.NewDecoder(resp.Body).Decode(fullResponse)

	resp.Body.Close()
	if err != nil {
		log.Println("Could not parse Poloniex json: ", err)
		return
	}

	btcTicker := (*fullResponse)["USDT_BTC"]
	ch <- btcTicker.GetExchangeData()
}

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
