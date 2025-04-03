package main

import (
	"pokedexcli/internal/pokecache"
)

const AreasEndpointURL = "https://pokeapi.co/api/v2/location-area/"

type cliCommand struct {
	name        string
	description string
	callback    func(confg *Config, cache *pokecache.Cache)
}

type Config struct {
	params string
	next_url string
	previous_url string
}

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var commandMapping = map[string]cliCommand{
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
	}
}

