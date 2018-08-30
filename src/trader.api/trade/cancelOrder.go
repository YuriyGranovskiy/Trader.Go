package trade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CancelOrderResponse struct {
	Success int               `json:"success"`
	Error   string            `json:"error"`
	Return  CancelOrderReturn `json:"return"`
}

type CancelOrderReturn struct {
	OrderId int                `json:"order_id"`
	Funds   map[string]float32 `json:"funds"`
}

func CancelOrder(orderId int, getRequest func(string, string, []byte) *http.Request, getNonce func() int) CancelOrderResponse {
	httpClient := CreateHttpClient()
	nonce := getNonce()
	requestBody := fmt.Sprintf("method=CancelOrder&nonce=%d&order_id=%d", nonce, orderId)
	request := getRequest(tradeApiUri, http.MethodPost, []byte(requestBody))

	res, getErr := httpClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	cancelOrderResponse := CancelOrderResponse{}
	jsonErr := json.Unmarshal(body, &cancelOrderResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return cancelOrderResponse
}
