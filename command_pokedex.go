package main

import (
	"fmt"
	"pokedexcli/internal/pokecache"
)

func commandPokedex(config *Config, cache *pokecache.Cache, myPokedex *Pokedex) {

	for pokemonName := range *myPokedex {
		fmt.Println("  -", pokemonName)
	}
	return
}
