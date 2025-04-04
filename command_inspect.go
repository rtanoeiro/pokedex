package main

import (
	"fmt"
	"pokedexcli/internal/pokecache"
)

func commandInspect(config *Config, cache *pokecache.Cache, myPokedex *Pokedex) {

	pokemonName := config.params
	ok := checkPokedex(pokemonName, myPokedex)

	if !ok {
		fmt.Println("You haven't caught this pokemon yet.")
		return
	}

	pokemonStats := (*myPokedex)[pokemonName]

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", pokemonStats.Height)
	fmt.Printf("Weight: %d\n", pokemonStats.Weight)
	fmt.Printf("Stats:\n")
	fmt.Printf("  - hp: %d\n", pokemonStats.Stats.hp)
	fmt.Printf("  - attack: %d\n", pokemonStats.Stats.attack)
	fmt.Printf("  - defense: %d\n", pokemonStats.Stats.defense)
	fmt.Printf("  - special-attack: %d\n", pokemonStats.Stats.specialAttack)
	fmt.Printf("  - special-defense: %d\n", pokemonStats.Stats.specialDefense)
	fmt.Printf("  - speed: %d\n", pokemonStats.Stats.speed)
	fmt.Printf("Types:\n")
	fmt.Printf("  - %s\n", pokemonStats.Types.type1)
	if pokemonStats.Types.type2 != "" {
		fmt.Printf("  - %s\n", pokemonStats.Types.type2)
	}
}
