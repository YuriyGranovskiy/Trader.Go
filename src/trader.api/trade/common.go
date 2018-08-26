package trade

import (
	"net/http"
	"net/url"
)

const tradeApiUri = "https://wex.nz/tapi"

func CreateHttpClient() http.Client {
	proxyUrl, _ := url.Parse("http://127.0.0.1:8888")
	return http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}
