package spotify_api

import (
	"net/http"
	"time"
)

const sp_baseURL = "https://api.spotify.com/v1"

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
