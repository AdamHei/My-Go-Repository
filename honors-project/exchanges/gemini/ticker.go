package gemini

import (
	"encoding/json"
	"github.com/adamhei/honors-project/exchanges/errorhandling"
	"github.com/adamhei/honors-project/exchanges/models"
	"net/http"
)

const tickerUrl = "https://api.gemini.com/v1/pubticker/btcusd"

func FetchBidAskData(ch chan<- models.LimitedJson) {
	resp, err := http.Get(tickerUrl)
	if err != nil {
		errorhandling.ErrorHandler("Could not fetch Gemini data:"+err.Error(), ch)
	}

	geminiResponse := new(models.GeminiTicker)
	err = json.NewDecoder(resp.Body).Decode(geminiResponse)

	resp.Body.Close()
	if err != nil {
		errorhandling.ErrorHandler("Could not parse Gemini json:"+err.Error(), ch)
		return
	}

	ch <- geminiResponse.GetExchangeData()
}
