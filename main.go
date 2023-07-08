package main

import (
	"fmt"
)

var (
	input string
)

func main() {



	for{
		fmt.Print("pokedex > ")

		fmt.Scanln(&input)

		if input == "help"{
			fmt.Println(`You should have two available commands:

			help: prints a help message describing how to use the REPL
			exit: exits the program`)

		}else{
			break
		}
	}
	

}
