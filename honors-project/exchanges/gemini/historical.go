package gemini

import (
	"encoding/json"
	"fmt"
	"github.com/adamhei/honors-project/exchanges/models"
	"net/http"
	"time"
)

const historyUrl = "https://api.gemini.com/v1/trades/btcusd?since=%d&limit_trades=500"

func GetTradeHistory(from time.Time, _ time.Time) []models.GeminiOrder {
	formattedUrl := fmt.Sprintf(historyUrl, from.Unix())
	resp, err := http.Get(formattedUrl)
	if err != nil {
		fmt.Println(err)
		return make([]models.GeminiOrder, 0)
	}

	orders := make([]models.GeminiOrder, 0)
	err = json.NewDecoder(resp.Body).Decode(&orders)
	resp.Body.Close()

	if err != nil {
		fmt.Println(err)
		return make([]models.GeminiOrder, 0)
	}

	return orders
}
