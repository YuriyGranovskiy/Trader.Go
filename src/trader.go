package main

import (
	"fmt"

	"trader.api/public"
	"trader.api/trade"
)

func main() {
	pairName := "btc_usd"
	infoResponse := public.GetInfo()
	fmt.Println(infoResponse.Pairs[pairName])
	depthResponse := public.GetDepthByPair(pairName)
	fmt.Println(depthResponse[pairName])
	activeOrdersResponse := trade.GetActiveOrdersByPair("btc_usd", trade.GetAuthenticatedRequest)
	fmt.Println(activeOrdersResponse)
	tradeHistoryResponse := trade.GetTradeHistory(trade.GetAuthenticatedRequest)
	fmt.Println(tradeHistoryResponse)
}
