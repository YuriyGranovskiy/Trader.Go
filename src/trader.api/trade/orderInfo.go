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

type OrderInfoResponse struct {
	Success int    `json:"success"`
	Error   string `json:"error"`
}

func GetOrderInfo(orderId int) OrderInfoResponse {
	proxyUrl, err := url.Parse("http://127.0.0.1:8888")
	httpClient := http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	nonce := 2
	requestBody := fmt.Sprintf("method=OrderInfo&nonce=%d&order_id=%d", nonce, orderId)
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

	orderInfoResponse := OrderInfoResponse{}
	jsonErr := json.Unmarshal(body, &orderInfoResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return orderInfoResponse
}
