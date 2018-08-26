package public

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const depthUrl = "https://wex.nz/api/3/depth"

type DepthPair struct {
	Asks []AskOrBid `json:"asks"`
	Bids []AskOrBid `json:"bids"`
}

type AskOrBid [2]float32

type DepthResponse map[string]DepthPair

func GetDepthByPair(pairName string) DepthResponse {
	proxyUrl, err := url.Parse("http://127.0.0.1:8888")
	httpClient := http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", depthUrl, pairName), nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	depthResponse := DepthResponse{}
	jsonErr := json.Unmarshal(body, &depthResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return depthResponse
}
