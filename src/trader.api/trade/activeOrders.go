package trade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ActiveOrdersResponse struct {
	Success int                 `json:"success"`
	Error   string              `json:"error"`
	Return  map[int]ActiveOrder `json:"return"`
}

type ActiveOrder struct {
	Pair      string  `json:"pair"`
	OrderType string  `json:"type"`
	Amount    float32 `json:"amount"`
	Rate      float32 `json:"rate"`
	Timestamp int     `json:"timestamp_created"`
	Status    int     `json:"status"`
}

func GetActiveOrdersByPair(pairName string, getRequest func(string, string, []byte) *http.Request, getNonce func() int) ActiveOrdersResponse {
	httpClient := CreateHttpClient()
	nonce := getNonce()
	requestBody := fmt.Sprintf("method=ActiveOrders&nonce=%d&pair=%s", nonce, pairName)
	request := getRequest(tradeApiUri, http.MethodPost, []byte(requestBody))

	res, getErr := httpClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	activeOrdersResponse := ActiveOrdersResponse{}
	jsonErr := json.Unmarshal(body, &activeOrdersResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return activeOrdersResponse
}
