package main

import (
	"encoding/json"
	"fmt"
	"pokedexcli/internal/pokecache"
)


func commandExplore(config *Config, cache *pokecache.Cache) {
	var httpErr error

	url := AreasEndpointURL + config.params

	resData, cacheError := cache.Get(url)
	
	if !cacheError {
		resData, httpErr = getRequest(url)

		if httpErr != nil {
			fmt.Println("Failed to get data from URL, try again")
			return
		}
	} 
	cache.Add(url, resData)

	explore := Explore{}	
	errUM := json.Unmarshal(resData, &explore)

	for _, pokemons := range explore.PokemonEncounters {
		fmt.Println(pokemons.Pokemon.Name)
	}
	if errUM != nil {
		fmt.Println("Got error Unmarshling Data")
		return
	}	
}