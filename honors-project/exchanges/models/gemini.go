package models

func (response GeminiTicker) GetExchangeData() LimitedJson {
	return LimitedJson{
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

type GeminiOrder struct {
	Timestamp   int64  `json:"timestamp"`
	TimestampMs int64  `json:"timestampms"`
	TID         int    `json:"tid"`
	Price       string `json:"price"`
	Amount      string `json:"amount"`
	Exchange    string `json:"exchange"`
	Type        string `json:"type"`
}
