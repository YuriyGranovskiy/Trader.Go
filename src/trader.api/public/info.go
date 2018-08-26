package public

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const infoUrl = "https://wex.nz/api/3/info"

type InfoResponse struct {
	ServerTime int                 `json:"server_time"`
	Pairs      map[string]InfoPair `json:"pairs"`
}

type InfoPair struct {
	DecimalPlaces int     `json:"decimal_places"`
	MinPrice      float32 `json:"min_price"`
	MaxPrice      float32 `json:"max_price"`
	MinAmount     float32 `json:"min_amount"`
	IsHidden      int     `json:"hidden"`
	Fee           float32 `json:"fee"`
}

func GetInfo() InfoResponse {
	httpClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, infoUrl, nil)
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
