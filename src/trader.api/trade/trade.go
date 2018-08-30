package trade

import (
	"fmt"
	"net/http"
)

type TradeResponse struct {
	Success int         `json:"success"`
	Error   string      `json:"error"`
	Return  TradeReturn `json:"return"`
}

type TradeReturn struct {
	Received float32            `json:"received"`
	Remains  float32            `json:"remains"`
	OrderId  int                `json:"order_id"`
	Funds    map[string]float32 `json:"funds"`
}

func Trade(pairName string, tradeType string, rate float32, amount float32, getRequest func(string, string, []byte) *http.Request, getNonce func() int) TradeResponse {
	httpClient := CreateHttpClient()
	nonce := getNonce()
	requestBody := fmt.Sprintf("method=Trade&nonce=%d&pair=%s&type=%s&rate=%g&amount=%g", nonce, pairName, tradeType, rate, amount)
	request := getRequest(tradeApiUri, http.MethodPost, []byte(requestBody))

	tradeResponse := TradeResponse{}

	ExecuteTradeRequest(httpClient, request, &tradeResponse)

	return tradeResponse
}
