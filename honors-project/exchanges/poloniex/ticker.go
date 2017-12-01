package poloniex

import (
	"encoding/json"
	"github.com/adamhei/honors-project/exchanges/errorhandling"
	"github.com/adamhei/honors-project/exchanges/models"
	"net/http"
)

const tickerUrl = "http://poloniex.com/public?command=returnTicker"
const tickerTag = "USDT_BTC"

func FetchBidAskPoloniex(ch chan<- models.LimitedJson) {
	resp, err := http.Get(tickerUrl)
	if err != nil {
		errorhandling.ErrorHandler("Could not fetch Poloniex data: "+err.Error(), ch)
		return
	}

	fullResponse := new(map[string]models.PoloniexTicker)
	err = json.NewDecoder(resp.Body).Decode(fullResponse)
	resp.Body.Close()

	if err != nil {
		errorhandling.ErrorHandler("Could not parse Poloniex json: "+err.Error(), ch)
		return
	}

	btcTicker := (*fullResponse)[tickerTag]
	ch <- btcTicker.GetExchangeData()
}
