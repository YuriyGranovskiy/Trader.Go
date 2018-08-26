package main

import (
	"fmt"

	"trader.api/public"
)

func main() {
	pairName := "btc_usd"
	infoResponse := public.GetInfo()
	fmt.Println(infoResponse.Pairs[pairName])
	depthResponse := public.GetDepthByPair(pairName)
	fmt.Println(depthResponse[pairName])
	/*tradeHistoryResponse := trade.GetTradeHistory()
	fmt.Println(tradeHistoryResponse)
	tradeHistoryResponse := trade.GetTradeHistory()
	fmt.Println(tradeHistoryResponse)*/
}
