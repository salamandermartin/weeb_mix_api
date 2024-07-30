package spotify_api

import (
	"net/http"
	"time"
)

const sp_baseURL = "https://api.spotify.com/v1"

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
