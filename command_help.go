package main

import (
	"fmt"
	"pokedexcli/internal/pokecache"
)

func getCommands()  map[string]cliCommand {

	commandMapping := map[string]cliCommand{
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
		"map":{
			name: "map",
			description: "Get the next 20 locations of the map",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Get the previous 20 locations of the map",
			callback: commandMapBack,
		},
		"explore": {
			name: "explore",
			description: "Explore the area to find available pokemons",
			callback: commandExplore,
		},
	}

	return commandMapping
}

func commandHelp(config *Config, cache *pokecache.Cache) {

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("Command %s: %s\n", command.name, command.description)
	}
	
}