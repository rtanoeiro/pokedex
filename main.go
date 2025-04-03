package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"pokedexcli/internal/pokecache"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(confg *Config, cache *pokecache.Cache)
}

type Config struct {
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

const AreasEndpointURL = "https://pokeapi.co/api/v2/location-area/"


func main() {
	commandMap := map[string]cliCommand{
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
	}

	cacheData := pokecache.NewCache(20*time.Second)
	config:= &Config{
		next_url: "",
		previous_url: "",
	}

	inputBuffer := bufio.NewScanner(os.Stdin)
	for {
	fmt.Print("Pokedex > ")
		if inputBuffer.Scan() {
			command, err := commandMap[inputBuffer.Text()]
			if !err  {
				fmt.Println("Unknown command, please try again")
				continue
			}
			command.callback(config, cacheData)
		}
	}
}

func commandExit(config *Config, cache *pokecache.Cache) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

func commandHelp(config *Config, cache *pokecache.Cache) {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Shows the next 20 locations of the map	`)
}

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