package main

import (
	"fmt"
	"os"
	"pokedexcli/internal/pokecache"
)

func commandExit(config *Config, cache *pokecache.Cache, myPokedex *Pokedex) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}
