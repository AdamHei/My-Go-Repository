package gdax

import (
	"encoding/json"
	"github.com/adamhei/honors-project/exchanges/errorhandling"
	"github.com/adamhei/honors-project/exchanges/models"
	"net/http"
)

const tickerUrl = "https://api.gdax.com/products/BTC-USD/ticker"

func FetchBidAskGDAX(ch chan<- models.LimitedJson) {
	resp, err := http.Get(tickerUrl)
	if err != nil {
		errorhandling.ErrorHandler("Could not fetch data from GDAX:"+err.Error(), ch)
		return
	}

	gdaxResponse := new(models.GDAXTicker)
	err = json.NewDecoder(resp.Body).Decode(gdaxResponse)
	resp.Body.Close()

	if err != nil {
		errorhandling.ErrorHandler("Could not parse GDAX json:"+err.Error(), ch)
		return
	}

	ch <- gdaxResponse.GetExchangeData()
}
