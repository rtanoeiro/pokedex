package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(confg *Config)
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
	}

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
			command.callback(config)
		}
	}
}

func commandExit(config *Config) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

func commandHelp(config *Config) {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Shows the next 20 locations of the map	`)
}

func commandMap(config *Config) {
	
	response, err := http.Get(AreasEndpointURL)

	if response.StatusCode != 200 {
		fmt.Println("Unable to get data, try again.")
		return
	}

	if err != nil {
		fmt.Println("Error getting data, try again.")
		return
	}
	locations := Locations{}

	resData, ioError := io.ReadAll(response.Body)
	if ioError != nil {
		fmt.Println("Found error when reading Body from HTTP Get Response")
		return
	}
	
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

func cleanInput(text string) []string {
	
	lowerWord := strings.ToLower(text)
	splitText := strings.Fields(lowerWord)

	return splitText
}