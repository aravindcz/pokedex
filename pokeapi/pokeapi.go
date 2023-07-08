package pokeapi

import (
	"errors"
	"encoding/json"
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

func GetPokeApiResult(command string,config *Config) (*Response,error){

	var request string
	var responseBodyUnmarshalled Response


	if config.Previous == "" && config.Next == ""{
		config.Next = "https://pokeapi.co/api/v2/location/?offset=0&limit=20"
	}

	if command == "map"{
		request = config.Next
	}else{
		if config.Previous == "null" || config.Previous == ""{
			
			return nil,errors.New("No previous link to follow")
		}
		request = config.Previous
	}

	response,err := http.Get(request)


	if err != nil {
		
		return nil,errors.New("Network error")
	}

	responseBody,err := io.ReadAll(response.Body)

	if err!=nil{
		
		return nil,errors.New("Parse error")
	}

	err = json.Unmarshal(responseBody,&responseBodyUnmarshalled)

	if err != nil{
		
		return nil,errors.New("Unmarshal error")
	}

	
	
	config.Next = responseBodyUnmarshalled.Next
	config.Previous = responseBodyUnmarshalled.Previous

	

	return &responseBodyUnmarshalled,nil

}

