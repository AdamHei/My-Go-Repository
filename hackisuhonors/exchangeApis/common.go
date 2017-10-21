package exchangeApis

type Ticker interface {
	GetBidAsk() []string
}
