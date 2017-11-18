package kraken

import (
	"github.com/adamhei/honors-project/exchanges/errorhandling"
	"github.com/adamhei/honors-project/exchanges/models"
	"net/http"
	"encoding/json"
)

const tickerUrl = "https://api.kraken.com/0/public/Ticker?pair=XBTUSD"

func FetchBidAskKraken(ch chan<- models.LimitedJson) {
	resp, err := http.Get(tickerUrl)
	if err != nil {
		errorhandling.ErrorHandler("Could not fetch Kraken data:"+err.Error(), ch)
	}

	krakenResponse := new(models.KrakenTicker)
	err = json.NewDecoder(resp.Body).Decode(krakenResponse)

	resp.Body.Close()
	if err != nil {
		errorhandling.ErrorHandler("Could not parse Kraken json:"+err.Error(), ch)
		return
	}

	ch <- krakenResponse.GetExchangeData()
}
