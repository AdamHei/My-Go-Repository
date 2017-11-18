package models

func (response KrakenTicker) GetExchangeData() LimitedJson {
	return LimitedJson{
		"Kraken": {
			"Bid": response.Result.XBTUSD.BidArr[0],
			"Ask": response.Result.XBTUSD.AskArr[0],
		},
	}
}

type KrakenTicker struct {
	Error  []string `json:"error"`
	Result result   `json:"result"`
}

type result struct {
	XBTUSD pairBody `json:"XXBTZUSD"`
}

type pairBody struct {
	AskArr               []string `json:"a"`
	BidArr               []string `json:"b"`
	ClosedTradeArr       []string `json:"c"`
	VolumeArr            []string `json:"v"`
	VolumeWeightedAvgArr []string `json:"p"`
	NumTradesArr         []int    `json:"t"`
	LowArr               []string `json:"l"`
	HighArr              []string `json:"h"`
	OpenPrice            string   `json:"o"`
}
