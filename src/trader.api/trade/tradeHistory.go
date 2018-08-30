package trade

import (
	"fmt"
	"net/http"
)

type TradeHistoryResponse struct {
	Success int                        `json:"success"`
	Error   string                     `json:"error"`
	Return  map[int]TradeHistoryReturn `json:"return"`
}

type TradeHistoryReturn struct {
	Pair        string  `json:"pair"`
	Type        string  `json:"type"`
	Amount      float32 `json:"amount"`
	Rate        float32 `json:"rate"`
	OrderId     int     `json:"order_id"`
	IsYourOrder int     `json:"is_your_order"`
	Timestamp   int     `json:"timestamp"`
}

func GetTradeHistory(getRequest func(string, string, []byte) *http.Request, getNonce func() int) TradeHistoryResponse {
	httpClient := CreateHttpClient()
	nonce := getNonce()
	requestBody := fmt.Sprintf("method=TradeHistory&nonce=%d&pair=btc_usd", nonce)

	request := getRequest(tradeApiUri, http.MethodPost, []byte(requestBody))

	tradeHistoryResponse := TradeHistoryResponse{}

	ExecuteTradeRequest(httpClient, request, &tradeHistoryResponse)

	return tradeHistoryResponse
}
