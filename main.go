package main

import (
	"fmt"
	"github.com/aravindcz/pokedex/internal/pokeapi"
	"github.com/aravindcz/pokedex/internal/pokecache"
)

var (
	input    string
	config   pokeapi.Config
	response string
	err      error
	cacheMap pokecache.CacheMap
)

func main() {

	config := pokeapi.Config{}

	cacheMap := pokecache.CacheMap{}
	

loop:
	for {
		fmt.Print("pokedex > ")

		fmt.Scanln(&input)

		switch input {

		case "help":

			fmt.Println(`You should have two available commands:

			help: prints a help message describing how to use the REPL
			exit: exits the program`)

		case "map", "mapb":
			response, err = pokeapi.GetPokeApiResult(input, &config, &cacheMap)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(response)
			}

		default:
			break loop

		}

	}

}
