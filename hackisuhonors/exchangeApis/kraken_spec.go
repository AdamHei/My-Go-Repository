package exchangeApis

const KRAKENURL = "https://api.kraken.com/0/public/Ticker?pair=XBTUSD"

func (response KrakenResponse) GetExchangeData() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"Kraken": {
			"Bid": response.Result.XBTUSD.BidArr[0],
			"Ask": response.Result.XBTUSD.AskArr[0],
		},
	}
}

type KrakenResponse struct {
	Error  []string `json:"error"`
	Result Result   `json:"result"`
}

type Result struct {
	XBTUSD PairBody `json:"XXBTZUSD"`
}

type PairBody struct {
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
