package exchangeApis

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
