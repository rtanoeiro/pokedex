package main

import (
	"fmt"
	"pokedexcli/internal/pokecache"
)

func commandHelp(config *Config, cache *pokecache.Cache) {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Shows the next 20 locations of the map	`)
}