package main

import (
	"fmt"

	"github.com/booorek/pokedexcli/internal/pokeAPI"
	"github.com/booorek/pokedexcli/internal/pokecache"
)
func commandMap(cache *pokecache.Cache,config *config) error {
	var gameMap pokeAPI.Locations
	var err error

	if config.next != "" {
		gameMap, err = pokeAPI.GetMapFromAPI(cache,&config.next)
	} else {
		gameMap, err = pokeAPI.GetMapFromAPI(nil)
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

func commandMapB(config *config) error {
	var gameMap pokeAPI.Locations
	var err error
	if config.previous != "" {
		gameMap, err = pokeAPI.GetMapFromAPI(&config.previous)
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


