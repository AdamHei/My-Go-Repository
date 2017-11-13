package exchanges

import "log"

const NUMEXCHANGES = 5

type Ticker interface {
	GetExchangeData() LimitedJson
}

type LimitedJson map[string]map[string]string

type MyError struct {
	Err     string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Err
}

func (e *MyError) ErrorCode() int {
	return e.ErrCode
}

// Print the error message and send an empty response through the channel
func ErrorHandler(errorMsg string, ch chan<- LimitedJson) {
	log.Print(errorMsg)
	ch <- make(LimitedJson)
}

func FetchAllExchanges(ch chan<- LimitedJson) {
	go fetchBidAskPoloniex(ch)
	go fetchBidAskGemini(ch)
	go fetchBidAskKraken(ch)
	go fetchBidAskGDAX(ch)
	go fetchBidAskBitfinex(ch)
}
