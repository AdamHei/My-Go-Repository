package exchangeApis

const GEMINIURL = "https://api.gemini.com/v1/pubticker/btcusd"

func (response GeminiTicker) GetBidAsk() []string {
	return []string{"Gemini", response.Bid, response.Ask}
}

type GeminiTicker struct {
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	BTCVolume volume `json:"volume"`
	Last      string `json:"last"`
}

type volume struct {
	BTC       string `json:"BTC"`
	USD       string `json:"USD"`
	Timestamp int    `json:"timestamp"`
}
