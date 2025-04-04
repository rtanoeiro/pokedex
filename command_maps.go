package main

import (
	"encoding/json"
	"fmt"
	"pokedexcli/internal/pokecache"
)

func commandMap(config *Config, cache *pokecache.Cache, myPokedex *Pokedex) {
	var url string
	var resData []byte
	var httpErr error

	if config.next_url == "" {
		url = AreasEndpointURL
	} else {
		url = config.next_url
	}

	resData, cacheError := cache.Get(url)
	
	/*Move this http request to a common func*/
	if !cacheError {
		resData, httpErr = getRequest(url)

		if httpErr != nil {
			fmt.Println("Failed to get data from URL, try again")
			return
		}
	} 

	locations := Locations{}	
	errUM := json.Unmarshal(resData, &locations)
	
	config.previous_url = locations.Previous
	config.next_url = locations.Next

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	if errUM != nil {
		fmt.Println("Got error Unmarshling Data")
		return
	}
}

func commandMapBack(config *Config, cache *pokecache.Cache, myPokedex *Pokedex) {
	var httpErr error

	if config.previous_url == "" {
		fmt.Println("you're in the first page")
		return
	}

	resData, cacheError := cache.Get(config.previous_url)
	
	if !cacheError {
		resData, httpErr = getRequest(config.previous_url)

		if httpErr != nil {
			fmt.Println("Failed to get data from URL, try again")
			return
		}
	} 
	locations := Locations{}

	errUM := json.Unmarshal(resData, &locations)
	
	config.previous_url = locations.Previous
	config.next_url = locations.Next

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	if errUM != nil {
		fmt.Println("Got error Unmarshling Data")
		return
	}
}