package main

import (
	"fmt"
	"pokeapi"
)


var (
	input string
	response pokeapi.Response
)

func main() {

	for{
		fmt.Print("pokedex > ")

		fmt.Scanln(&input)

		if input == "help"{
			fmt.Println(`You should have two available commands:

			help: prints a help message describing how to use the REPL
			exit: exits the program`)

		}else if input== "map" || input == "mapb"{
			 response  = pokeapi.GetPokeApiResult(input)
			 fmt.Println(string(response))
		}else{
			break
		}
	}

}
