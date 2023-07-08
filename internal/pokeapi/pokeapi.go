package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
	"github.com/aravindcz/pokedex/internal/pokecache"
)

type Response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Previous string
	Next     string
}

func GetPokeApiResult(command string, config *Config, cacheMap *pokecache.CacheMap) (string, error) {

	var request string
	var responseBodyUnmarshalled Response
	var response *http.Response
	var responseBody []byte
	var responseString string

	if config.Previous == "" && config.Next == "" {
		cacheMap.NewCache(time.Duration(5000))
		config.Next = "https://pokeapi.co/api/v2/location/?offset=0&limit=20"
	}

	if command == "map" {
		request = config.Next
	} else {
		if config.Previous == "null" || config.Previous == "" {

			return "", errors.New("No previous link to follow")
		}
		request = config.Previous
	}

	cacheResponse, err := cacheMap.Get(request)
	if err != nil {
		response, err = http.Get(request)
		if err != nil {

			return "", errors.New("Network error")
		}

		responseBody, err = io.ReadAll(response.Body)

		if err != nil {

			return "", errors.New("Parse error")
		}

		err = json.Unmarshal(responseBody, &responseBodyUnmarshalled)

		if err != nil {

			return "", errors.New("Unmarshal error")
		}

		cacheMap.Add(request, string(responseBody))
		responseString = string(responseBody)

	} else {
		responseString = cacheResponse.Value
	}

	config.Next = responseBodyUnmarshalled.Next
	config.Previous = responseBodyUnmarshalled.Previous
	

	return responseString, nil

}
