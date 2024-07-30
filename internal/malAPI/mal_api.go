package mal_api

import (
	"net/http"
	"time"
)

const mal_baseURL = "https://api.myanimelist.net/v2"

type mal_Client struct {
	httpClient http.Client
}

func malNewClient(cacheInterval time.Duration) mal_Client {
	return mal_Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
