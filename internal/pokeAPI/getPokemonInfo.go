package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(url *string) (Pokemon, error) {
	var finalPath string
	if url == nil {
		return Pokemon{}, fmt.Errorf("Cannot get info without name")
	}

	finalPath = apiURL + "/pokemon/" + *url
	if res, ok := c.cache.Get(finalPath); ok {
		data := Pokemon{}
		err := json.Unmarshal(res, &data)
		if err != nil {
			return Pokemon{}, err
		}
		return data, nil
	}

	res, err := http.Get(finalPath)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error while acquiring data from API")
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("Error while reading from body response")
	}
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return Pokemon{}, fmt.Errorf("Error while reading from body response")
	}

	data := Pokemon{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		return Pokemon{}, fmt.Errorf("Cannot unmarshal data")
	}
	c.cache.Add(finalPath, body)

	return data, nil
}
