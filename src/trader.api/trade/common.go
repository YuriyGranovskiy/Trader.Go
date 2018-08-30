package trade

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const tradeApiUri = "https://wex.nz/tapi"

func CreateHttpClient() http.Client {
	proxyUrl, _ := url.Parse("http://127.0.0.1:8888")
	return http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}

func ExecuteTradeRequest(httpClient http.Client, request *http.Request, response interface{}) {
	res, getErr := httpClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}
