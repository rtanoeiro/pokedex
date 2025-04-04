package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

func main() {
	cacheData := pokecache.NewCache(20*time.Second)
	config:= &Config{
		params: "",
		next_url: "",
		previous_url: "",
	}

	inputBuffer := bufio.NewScanner(os.Stdin)
	for {
	fmt.Print("Pokedex > ")
		if inputBuffer.Scan() {
			
			commands := getCommands()
			parameters := strings.Fields(inputBuffer.Text())
			command, err := commands[parameters[0]]
			
			if len(parameters) > 1 {
				config.params = parameters[1]
			}

			if !err  {
				fmt.Println("Unknown command, please try again")
				continue
			}

			command.callback(config, cacheData)
		}
	}
}