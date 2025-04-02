package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}


func main() {
	commandMap := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Contains usage details for the CLI",
			callback: commandHelp,
		},
	}

	inputBuffer := bufio.NewScanner(os.Stdin)
	for {
	fmt.Print("Pokedex > ")
		if inputBuffer.Scan() {
			command, err := commandMap[inputBuffer.Text()]
			fmt.Println("Used command: ", command)
			fmt.Println("Currenf error: ", err)
			if !err  {
				fmt.Println("Unknown command, please try again")
				continue
			}
			command.callback()
		}
	}
}

func commandExit() {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

func commandHelp() {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
}

func cleanInput(text string) []string {
	
	lowerWord := strings.ToLower(text)
	splitText := strings.Fields(lowerWord)

	return splitText
}