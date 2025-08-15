package main 

import(
	"github.com/booorek/pokedexcli/internal/pokeAPI"
	"fmt"
)
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


