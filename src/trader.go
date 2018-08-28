package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"trader.api/public"
	"trader.api/trade"
)

type Config struct {
	KeyFilePath   string `json:"keyFile"`
	NonceFilePath string `json:"nonceFile"`
	PairName      string `json:"pairName"`
}

var nonce int

func main() {
	config := loadConfiguration("config.json")
	nonce = getNonceFromFile(normalizePath(config.NonceFilePath))
	infoResponse := public.GetInfo()
	fmt.Println(infoResponse.Pairs[config.PairName])
	depthResponse := public.GetDepthByPair(config.PairName)
	fmt.Println(depthResponse[config.PairName])
	activeOrdersResponse := trade.GetActiveOrdersByPair(config.PairName, requestFactory(normalizePath(config.KeyFilePath)), func() int {
		return getNonce(normalizePath(config.NonceFilePath))
	})
	fmt.Println(activeOrdersResponse)
	tradeHistoryResponse := trade.GetTradeHistory(requestFactory(normalizePath(config.KeyFilePath)), func() int {
		return getNonce(normalizePath(config.NonceFilePath))
	})
	fmt.Println(tradeHistoryResponse)
}

func loadConfiguration(configFilePath string) Config {
	file, e := ioutil.ReadFile(configFilePath)
	if e != nil {
		log.Fatalf("Can't load file. %s", e)
	}

	config := Config{}
	err := json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Can't parse config file. %s", err)
	}

	return config
}

func requestFactory(keyFilePath string) func(string, string, []byte) *http.Request {
	file, e := ioutil.ReadFile(keyFilePath)
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

func getNonceFromFile(nonceFilePath string) int {

	file, e := ioutil.ReadFile(nonceFilePath)
	if e != nil {
		log.Fatalf("Can't load nonce file. %s", e)
	}

	nonceFromFile, e := strconv.Atoi(string(file))
	if e != nil {
		log.Fatalf("Nonce file is invalid. %s", e)
	}

	return nonceFromFile
}

func getNonce(nonceFilePath string) int {
	currentNonce := nonce
	nonce++
	ioutil.WriteFile(nonceFilePath, []byte(strconv.Itoa(nonce)), 0)
	return currentNonce
}

func normalizePath(path string) string {
	re := regexp.MustCompile("%\\S+%")
	envVar := re.FindString(path)
	if envVar == "" {
		return path
	}

	resolvedEnvVar := os.Getenv(strings.Trim(envVar, "%"))

	return strings.Replace(path, envVar, resolvedEnvVar, -1)
}
