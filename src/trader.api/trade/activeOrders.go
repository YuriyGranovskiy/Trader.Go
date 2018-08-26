package trade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type ActiveOrdersResponse struct {
	Success int    `json:"success"`
	Error   string `json:"error"`
}

func GetActiveOrdersByPair(pairName string, getRequest func(string, string, []byte) *http.Request) ActiveOrdersResponse {
	proxyUrl, _ := url.Parse("http://127.0.0.1:8888")
	httpClient := http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	nonce := 6
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
