package trade

import (
	"fmt"
	"net/http"
)

type OrderInfoResponse struct {
	Success int    `json:"success"`
	Error   string `json:"error"`
}

func GetOrderInfo(orderId int, getRequest func(string, string, []byte) *http.Request) OrderInfoResponse {
	httpClient := CreateHttpClient()
	nonce := 2
	requestBody := fmt.Sprintf("method=OrderInfo&nonce=%d&order_id=%d", nonce, orderId)

	request := getRequest(tradeApiUri, http.MethodPost, []byte(requestBody))

	orderInfoResponse := OrderInfoResponse{}

	ExecuteTradeRequest(httpClient, request, &orderInfoResponse)

	return orderInfoResponse
}
