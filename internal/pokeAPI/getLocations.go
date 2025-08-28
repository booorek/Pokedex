package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetMapFromAPI(url *string) (Locations, error) {
	var finalPath string
	if url == nil {
		finalPath = apiURL + "/location-area"
	} else {
		finalPath = *url
	}
	if res, ok := c.cache.Get(finalPath); ok {
		data := Locations{}
		err := json.Unmarshal(res, &data)
		if err != nil {
			return Locations{}, err
		}
		return data,nil
	}
	res, err := http.Get(finalPath)
	if err != nil {
		return Locations{}, fmt.Errorf("Error while acquiring data from API")
	}

	if res.StatusCode > 299 {
		return Locations{}, fmt.Errorf("Error while reading from body response")
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		return Locations{}, fmt.Errorf("Error while reading from body response")
	}

	data := Locations{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return Locations{}, fmt.Errorf("Cannot unmarshal data")
	}

	c.cache.Add(finalPath,body)
	return data, nil
}
