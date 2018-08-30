package trade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	res, getErr := httpClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	tradeResponse := TradeResponse{}
	jsonErr := json.Unmarshal(body, &tradeResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return tradeResponse
}
