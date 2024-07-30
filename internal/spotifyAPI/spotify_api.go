package spotify_api

import (
	"net/http"
	"time"
)

type spotify_Client struct {
	httpClient http.Client
}

func spotifyNewClient(cacheInterval time.Duration) spotify_Client {
	return spotify_Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
