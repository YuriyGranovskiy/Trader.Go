package trade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func GetTradeHistory(getRequest func(string, string, []byte) *http.Request) TradeHistoryResponse {
	proxyUrl, _ := url.Parse("http://127.0.0.1:8888")
	httpClient := http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	nonce := 7
	requestBody := fmt.Sprintf("method=TradeHistory&nonce=%d&pair=btc_usd", nonce)

	request := getRequest(tradeApiUri, http.MethodPost, []byte(requestBody))

	res, getErr := httpClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	tradeHistoryResponse := TradeHistoryResponse{}
	jsonErr := json.Unmarshal(body, &tradeHistoryResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return tradeHistoryResponse
}
