package spotify_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) SongReq(pageURL *string, songName string) (SearchQueryResp, error) {
	endpoint := "/search"
	fullURL := sp_baseURL + endpoint + songName

	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return SearchQueryResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return SearchQueryResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return SearchQueryResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return SearchQueryResp{}, err
	}

	searchQueryResp := SearchQueryResp{}
	err = json.Unmarshal(dat, &searchQueryResp)

	return searchQueryResp, nil
}
