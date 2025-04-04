package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"pokedexcli/internal/pokecache"
)
func commandCatch(config *Config, cache *pokecache.Cache, myPokedex *Pokedex) {

	pokemonName := config.params
	url:=PokemonURL+pokemonName
	resData, err := getRequest(url)

	if err != nil {
		fmt.Println("Unable to find Pokemon", pokemonName, ". Error:", err)
	}
	
	pokemons := Pokemon{}
	errUM := json.Unmarshal(resData, &pokemons)
	
	if errUM!= nil {
		fmt.Println("Error unmarshaling data: ", errUM)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	baseDifficulty := pokemons.BaseExperience
	catchRate := float64(baseDifficulty)/256.0
	randomNumber := rand.Float64()

	var exists bool
	if randomNumber >= catchRate {
		exists = checkPokedex(pokemonName, myPokedex)
		fmt.Printf("%s was caught!\n", pokemonName)
		if !exists {
			addToPokedex(pokemonName, myPokedex)
		} else {
			updatedPokedex(pokemonName, myPokedex)
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	
}

func checkPokedex(pokemon string, pokedex *Pokedex) bool {

	_, ok := (*pokedex)[pokemon]

	return ok
}

func addToPokedex(pokemon string, pokedex *Pokedex) {
	(*pokedex)[pokemon] = 1
}

func updatedPokedex(pokemon string, pokedex *Pokedex) {
	(*pokedex)[pokemon]++
}