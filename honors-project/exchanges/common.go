package exchanges

import (
	"github.com/adamhei/honors-project/exchanges/bitfinex"
	"github.com/adamhei/honors-project/exchanges/gdax"
	"github.com/adamhei/honors-project/exchanges/gemini"
	"github.com/adamhei/honors-project/exchanges/kraken"
	"github.com/adamhei/honors-project/exchanges/models"
	"github.com/adamhei/honors-project/exchanges/poloniex"
)

const NUMEXCHANGES = 5

func FetchAllExchanges(ch chan<- models.LimitedJson) {
	go poloniex.FetchBidAskPoloniex(ch)
	go gemini.FetchBidAskData(ch)
	go kraken.FetchBidAskKraken(ch)
	go gdax.FetchBidAskGDAX(ch)
	go bitfinex.FetchBidAskBitfinex(ch)
}
