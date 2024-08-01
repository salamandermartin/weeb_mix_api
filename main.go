package main

import (
	"time"

	mal_api "github.com/salamandermartin/weeb_mix_api/internal/malAPI"
	spotify_api "github.com/salamandermartin/weeb_mix_api/internal/spotifyAPI"
)

type config struct {
	spotifyClient spotify_api.Client
	malClient     mal_api.Client
}

func main() {
	cfg := config{
		spotifyClient: spotify_api.NewClient(time.Hour),
		malClient:     mal_api.NewClient(time.Hour),
	}
}
