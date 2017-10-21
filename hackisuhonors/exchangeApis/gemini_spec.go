package exchangeApis

const GEMINIURL = "https://api.gemini.com/v1/pubticker/btcusd"

func (response GeminiTicker) GetExchangeData() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"Gemini": {
			"Bid": response.Bid,
			"Ask": response.Ask,
		},
	}
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
