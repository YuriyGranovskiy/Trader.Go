package trade

import (
	"bytes"
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

func GetTradeHistory() TradeHistoryResponse {
	proxyUrl, err := url.Parse("http://127.0.0.1:8888")
	httpClient := http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	nonce := 3
	requestBody := fmt.Sprintf("method=TradeHistory&nonce=%d&pair=btc_usd", nonce)
	data := []byte(requestBody)
	r := bytes.NewReader(data)

	req, err := http.NewRequest(http.MethodPost, tradeApiUri, r)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Key", AuthKey)
	req.Header.Set("Sign", computeHmac512(requestBody, AuthSecret))

	res, getErr := httpClient.Do(req)
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
