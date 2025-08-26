package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Locations struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}
type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (c *Client) GetMapFromAPI(path *string) (Locations, error) {
	var finalPath string
	if path == nil {
		finalPath = "https://pokeapi.co/api/v2/location-area/"
	} else {
		finalPath = *path
	}
	res, err := http.Get(finalPath)
	if err != nil {
		return Locations{}, fmt.Errorf("Error while acquiring data from API")
	}
	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Locations{}, fmt.Errorf("Error while reading from body response")
	}
	if err != nil {
		return Locations{}, fmt.Errorf("Error while reading from body response")
	}

	data := Locations{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return Locations{}, fmt.Errorf("Cannot unmarshal data")
	}
	return data, nil
}
