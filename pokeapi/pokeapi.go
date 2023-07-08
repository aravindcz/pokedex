package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


	
type Response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct{
	Previous string
	Next string
}

func GetPokeApiResult(command string,config *Config) *Response{

	var request string
	var responseBodyUnmarshalled Response


	if config.Previous == ""{
		config.Next = "https://pokeapi.co/api/v2/location/?offset=0&limit=20"
	}

	if command == "map"{
		request = config.Next
	}else{
		if config.Previous == null || config.Previous == ""{
			return nil
		}
		request = config.Previous
	}

	response,err := http.Get(request)


	if err != nil {
		fmt.Println("Some error occured")
		return nil
	}

	responseBody,err := io.ReadAll(response.Body)

	if err!=nil{
		return nil
	}

	err = json.Unmarshal(responseBody,&responseBodyUnmarshalled)

	if err != nil{
		return nil
	}

	config.Next = responseBodyUnmarshalled.Next
	config.Previous = responseBodyUnmarshalled.Previous

	return &responseBodyUnmarshalled

}

