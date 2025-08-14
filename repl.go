package main

import (
	"bufio"
	"fmt"
	"github.com/booorek/pokedexcli/internal/pokeAPI"
	"os"
	"strings"
)

var commandRegistry map[string]cliCommand

func startPokedex() {
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
		if err := command.callback(cfg); err != nil{
			fmt.Printf("Error: %v\n",err)
		}

	}
}

func commandExit(config *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n")
	for _, command := range commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(config *config) error {
	var gameMap pokeAPI.Locations
	var err error

	if config.next != "" {
		gameMap, err = pokeAPI.GetMapFromAPI(&config.next)
	} else {
		gameMap, err = pokeAPI.GetMapFromAPI(nil)
	}
	if err != nil {
		fmt.Errorf("Error while communicating with API\n%v", err)
		return err
	}

	config.next = gameMap.Next
	config.previous = gameMap.Previous

	for _, location := range gameMap.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func commandMapB(config *config) error {
	var gameMap pokeAPI.Locations
	var err error
	if config.previous != "" {
		gameMap, err = pokeAPI.GetMapFromAPI(&config.previous)
	} else {
		return fmt.Errorf("No backward move\n")
	}
	if err != nil {
		fmt.Errorf("Error while communicating with API\n%v", err)
		return err
	}
	config.next = gameMap.Next
	config.previous = gameMap.Previous

	for _, location := range gameMap.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
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
