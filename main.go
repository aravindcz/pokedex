package main

import (
	"fmt"
	"github.com/aravindcz/pokedex/pokeapi"
)


var (
	input string
	config *pokeapi.Config
	response *pokeapi.Response
)

func main() {

	config := &pokeapi.Config{
		Next: ""
		Previous: ""
	}


	for{
		fmt.Print("pokedex > ")

		fmt.Scanln(&input)

		if input == "help"{
			fmt.Println(`You should have two available commands:

			help: prints a help message describing how to use the REPL
			exit: exits the program`)

		}else if input== "map" || input == "mapb"{
			 response  = pokeapi.GetPokeApiResult(input)
			 fmt.Println(*response)
		}else{
			break
		}
	}

}
