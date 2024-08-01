package main

import (
	"time"
)

type config struct {
	spotifyClient spotify_api.spotify_Client
	malClient     mal_api.mal_Client
}

func main() {
	cfg := config{
		spotifyClient: spotify_api.spotifyNewClient(time.Hour),
		malClient:     mal_api.malNewClient(time.Hour),
	}
}
