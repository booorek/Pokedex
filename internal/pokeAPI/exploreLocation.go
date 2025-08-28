package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(url *string) (AreaContent, error) {
	var finalPath string
	if url == nil {
		return AreaContent{}, fmt.Errorf("Cannot explore without argument")
	}

	finalPath = apiURL + "/location-area/" + *url
	if res, ok := c.cache.Get(finalPath); ok {
		data := AreaContent{}
		err := json.Unmarshal(res, &data)
		if err != nil {
			return AreaContent{}, err
		}
		return data, nil
	}

	res, err := http.Get(finalPath)
	if err != nil {
		return AreaContent{}, fmt.Errorf("Error while acquiring data from API")
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return AreaContent{}, fmt.Errorf("Error while reading from body response")
	}
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return AreaContent{}, fmt.Errorf("Error while reading from body response")
	}

	data := AreaContent{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		return AreaContent{}, fmt.Errorf("Cannot unmarshal data")
	}
	c.cache.Add(finalPath, body)

	return data, nil
}
