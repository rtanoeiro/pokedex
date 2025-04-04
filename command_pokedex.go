package main

import (
	"fmt"
	"pokedexcli/internal/pokecache"
)

func commandPokedex(config *Config, cache *pokecache.Cache, myPokedex *Pokedex) {

	for pokemonName, _ := range *myPokedex {
		fmt.Println("  -", pokemonName)
	}
	return
}
