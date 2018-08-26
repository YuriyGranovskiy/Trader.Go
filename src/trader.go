package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"trader.api/public"
	"trader.api/trade"
)

func main() {
	pairName := "btc_usd"
	infoResponse := public.GetInfo()
	fmt.Println(infoResponse.Pairs[pairName])
	depthResponse := public.GetDepthByPair(pairName)
	fmt.Println(depthResponse[pairName])
	activeOrdersResponse := trade.GetActiveOrdersByPair("btc_usd", RequestFactory())
	fmt.Println(activeOrdersResponse)
	tradeHistoryResponse := trade.GetTradeHistory(RequestFactory())
	fmt.Println(tradeHistoryResponse)
}

func RequestFactory() func(string, string, []byte) *http.Request {
	file, e := ioutil.ReadFile("../keys")
	if e != nil {
		log.Fatalf("Can't load file. %s", e)
	}

	fileParts := strings.Split(string(file), ":")

	if len(fileParts) != 2 {
		log.Fatal("Improper file structure")
	}

	key := fileParts[0]
	secret := fileParts[1]
	return func(uri string, method string, body []byte) *http.Request {
		return trade.GetAuthenticatedRequest(uri, method, body, key, secret)
	}
}
