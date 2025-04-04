package main

import (
	"fmt"
	"pokedexcli/internal/pokecache"
)

func commandHelp(config *Config, cache *pokecache.Cache, myPokedex *Pokedex) {

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("Command %s: %s\n", command.name, command.description)
	}
	
}