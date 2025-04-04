package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"pokedexcli/internal/pokecache"
)

func commandCatch(config *Config, cache *pokecache.Cache, myPokedex *Pokedex) {

	pokemonName := config.params
	url := PokemonURL + pokemonName
	resData, err := getRequest(url)

	if err != nil {
		fmt.Println("Unable to find Pokemon", pokemonName, ". Error:", err)
	}

	thisPokemon := Pokemon{}
	errUM := json.Unmarshal(resData, &thisPokemon)

	if errUM != nil {
		fmt.Println("Error unmarshaling data: ", errUM)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	baseDifficulty := thisPokemon.BaseExperience
	catchRate := float64(baseDifficulty) / 256.0
	randomNumber := rand.Float64()
	var exists bool
	if randomNumber >= catchRate {
		exists = checkPokedex(pokemonName, myPokedex)
		fmt.Printf("%s was caught!\n", pokemonName)
		if !exists {
			addToPokedex(thisPokemon, pokemonName, myPokedex)
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return
}

func checkPokedex(pokemonName string, pokedex *Pokedex) bool {

	_, ok := (*pokedex)[pokemonName]

	return ok
}

func addToPokedex(pokemonDetails Pokemon, pokemonName string, pokedex *Pokedex) {

	hp, atk, def, spatk, spdef, spd := getPokemonStats(pokemonDetails, pokemonName)
	types := getPokemonTypes(pokemonDetails, pokemonName)

	pokemon := (*pokedex)[pokemonName]
	pokemon.Height = pokemonDetails.Height
	pokemon.Weight = pokemonDetails.Weight
	pokemon.Stats.hp = hp
	pokemon.Stats.attack = atk
	pokemon.Stats.defense = def
	pokemon.Stats.specialAttack = spatk
	pokemon.Stats.specialDefense = spdef
	pokemon.Stats.speed = spd

	if len(types) == 2 {
		pokemon.Types.type1 = types[0]
		pokemon.Types.type2 = types[1]
	} else {
		pokemon.Types.type1 = types[0]
	}

	(*pokedex)[pokemonName] = pokemon

	return
}

func getPokemonStats(pokemonDetails Pokemon, pokemonName string) (int, int, int, int, int, int) {

	pokemonStats := pokemonDetails.Stats

	hp := pokemonStats[0].BaseStat
	atk := pokemonStats[1].BaseStat
	defense := pokemonStats[2].BaseStat
	spatk := pokemonStats[3].BaseStat
	spdef := pokemonStats[4].BaseStat
	spd := pokemonStats[5].BaseStat

	return hp, atk, defense, spatk, spdef, spd
}

func getPokemonTypes(pokemonDetails Pokemon, pokemonName string) []string {

	pokemonStats := pokemonDetails.Types

	var pokemonTypes []string

	for _, pokemonType := range pokemonStats {
		pokemonTypes = append(pokemonTypes, pokemonType.Type.Name)
	}

	return pokemonTypes
}
