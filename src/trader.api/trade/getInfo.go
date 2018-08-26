package trade

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type InfoResponse struct {
	Success int        `json:"success"`
	Return  InfoReturn `json:"return"`
}

type InfoReturn struct {
	Permissions      InfoPermissions `json:"rights"`
	Funds            InfoFunds       `json:"funds"`
	TransactionCount int             `json:"transaction_count"`
	OpenOrders       int             `json:"open_orders"`
	ServerTime       int             `json:"server_time"`
}

type InfoFunds struct {
	Btc float32 `json:"btc"`
	Usd float32 `json:"usd"`
}

type InfoPermissions struct {
	Info     int `json:"info"`
	Trade    int `json:"trade"`
	Withdraw int `json:"withdraw"`
}

func GetInfo() InfoResponse {
	httpClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, tradeApiUri, nil)
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

	infoResponse := InfoResponse{}
	jsonErr := json.Unmarshal(body, &infoResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return infoResponse
}
