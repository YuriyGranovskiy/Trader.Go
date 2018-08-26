package trade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	res, getErr := httpClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	orderInfoResponse := OrderInfoResponse{}
	jsonErr := json.Unmarshal(body, &orderInfoResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return orderInfoResponse
}
