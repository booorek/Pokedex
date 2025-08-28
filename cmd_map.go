package main

import (
	"fmt"
	"github.com/booorek/pokedexcli/internal/pokeAPI"
)

func commandMap(config *config, args []string) error {
	var gameMap pokeAPI.Locations
	var err error

	if config.next != "" {
		gameMap, err = config.pokeapiClient.GetMapFromAPI(&config.next)
	} else {
		gameMap, err = config.pokeapiClient.GetMapFromAPI(nil)

	}

	if err != nil {
		return err
	}

	config.next = gameMap.Next
	config.previous = gameMap.Previous

	for _, location := range gameMap.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func commandMapB(config *config, args []string) error {
	var gameMap pokeAPI.Locations
	var err error
	if config.previous != "" {
		gameMap, err = config.pokeapiClient.GetMapFromAPI(&config.previous)
	} else {
		return fmt.Errorf("No backward move\n")
	}
	if err != nil {
		return err
	}
	config.next = gameMap.Next
	config.previous = gameMap.Previous

	for _, location := range gameMap.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}
