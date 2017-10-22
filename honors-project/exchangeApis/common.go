package exchangeApis

const NUMEXCHANGES = 5

type Ticker interface {
	GetExchangeData() map[string]map[string]string
}

type MyError struct {
	Err       string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Err
}

func (e *MyError) ErrorCode() int {
	return e.ErrCode
}

func FetchAllExchanges(ch chan<- map[string]map[string]string) {
	go fetchBidAskPoloniex(ch)
	go fetchBidAskGemini(ch)
	go fetchBidAskKraken(ch)
	go fetchBidAskGDAX(ch)
	go fetchBidAskBitfinex(ch)
}