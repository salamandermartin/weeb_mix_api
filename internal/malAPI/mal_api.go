package mal_api

import (
	"net/http"
	"time"
)

const mal_baseURL = "https://api.myanimelist.net/v2"

type Client struct {
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
