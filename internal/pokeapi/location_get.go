package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) GetLocation(locationName string) (LocationData, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok { //Use this if cached
		locationsResp := LocationData{} //Still need to unmarshal
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationData{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationData{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationData{}, err
	}

	locationsResp := LocationData{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationData{}, err
	}

	c.cache.Add(url, dat) //Add new value to cache
	return locationsResp, nil
}
