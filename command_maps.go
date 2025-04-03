package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
)

func commandMap(config *Config, cache *pokecache.Cache) {
	var url string
	var resData []byte
	var ioError error

	if config.next_url == "" {
		url = AreasEndpointURL
	} else {
		url = config.next_url
	}

	resData, cacheError := cache.Get(url)
	
	/*Move this http request to a common func*/
	if !cacheError {
		response, httpErr := http.Get(url)
		
		if response.StatusCode != 200 {
			fmt.Println("Unable to get data, try again.")
			return
		}
	
		if httpErr != nil {
			fmt.Println("Error getting data, try again.")
			return
		}

		resData, ioError = io.ReadAll(response.Body)

		if ioError != nil {
			fmt.Println("Found error when reading Body from HTTP Get Response")
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

func commandMapBack(config *Config, cache *pokecache.Cache) {
	var ioError error

	if config.previous_url == "" {
		fmt.Println("you're in the first page")
		return
	}

	resData, cacheError := cache.Get(config.previous_url)
	
	if !cacheError {
		response, httpErr := http.Get(config.previous_url)
		
		if response.StatusCode != 200 {
			fmt.Println("Unable to get data, try again.")
			return
		}
	
		if httpErr != nil {
			fmt.Println("Error getting data, try again.")
			return
		}

		resData, ioError = io.ReadAll(response.Body)

		if ioError != nil {
			fmt.Println("Found error when reading Body from HTTP Get Response")
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