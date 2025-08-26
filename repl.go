package main

import (
	"bufio"
	"fmt"
	"github.com/booorek/pokedexcli/internal/pokeAPI"
	"os"
	"strings"
)

var commandRegistry map[string]cliCommand

func startPokedex(c *pokeAPI.Client) {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command, err := commandRegistry[string(scanner.Text())]
		if !err {
			fmt.Printf("Unknown command\n")
			continue
		}

		if err := command.callback(cfg); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func cleanInput(text string) []string {
	finalString := strings.Fields(strings.ToLower(text))
	return finalString
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
type config struct {
	next     string
	previous string
}

func init() {
	commandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Prints help messages",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows available locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows available back locations",
			callback:    commandMapB,
		},
	}
}
