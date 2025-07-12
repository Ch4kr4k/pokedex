package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreaRespStruct, error) {

	endpoint := "/location-area"
	fullUrl := baseURL + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	// check the cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache Hit")
		locatioAreasResp := LocationAreaRespStruct{}
		err := json.Unmarshal(dat, &locatioAreasResp)
		if err != nil {
			return LocationAreaRespStruct{}, err
		}
		return locatioAreasResp, nil
	}
	fmt.Println("Cache Miss")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaRespStruct{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaRespStruct{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaRespStruct{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaRespStruct{}, nil
	}

	locatioAreasResp := LocationAreaRespStruct{}
	err = json.Unmarshal(dat, &locatioAreasResp)
	if err != nil {
		return LocationAreaRespStruct{}, err
	}
	c.cache.Add(fullUrl, dat)
	return locatioAreasResp, nil
}
