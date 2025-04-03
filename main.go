package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokecache"
	"time"
)

func main() {
	cacheData := pokecache.NewCache(20*time.Second)
	config:= &Config{
		next_url: "",
		previous_url: "",
	}

	inputBuffer := bufio.NewScanner(os.Stdin)
	for {
	fmt.Print("Pokedex > ")
		if inputBuffer.Scan() {
			command, err := commandMapping[inputBuffer.Text()]
			if !err  {
				fmt.Println("Unknown command, please try again")
				continue
			}
			command.callback(config, cacheData)
		}
	}
}