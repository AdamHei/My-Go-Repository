package bitfinex

import (
	"encoding/json"
	"github.com/adamhei/honors-project/exchanges/errorhandling"
	"github.com/adamhei/honors-project/exchanges/models"
	"net/http"
)

const tickerUrl = "https://api.bitfinex.com/v1/pubticker/btcusd"

func FetchBidAskBitfinex(ch chan<- models.LimitedJson) {
	resp, err := http.Get(tickerUrl)
	if err != nil {
		errorhandling.ErrorHandler("Could not fetch data from Bitfinex:"+err.Error(), ch)
		return
	}

	bitResponse := new(models.BitfinexTicker)
	err = json.NewDecoder(resp.Body).Decode(bitResponse)

	resp.Body.Close()
	if err != nil {
		errorhandling.ErrorHandler("Could not parse Bitfinex response"+err.Error(), ch)
		return
	}

	ch <- bitResponse.GetExchangeData()
}
