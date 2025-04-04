package main


func getCommands()  map[string]cliCommand {

	commandMapping := map[string]cliCommand{
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
		},
		"catch": {
			name: "catch",
			description: "Attemps to catch a Pokemon by throwing a pokeball",
			callback: commandCatch,
		},
	}

	return commandMapping
}