package main

import (
	"fmt"

	"github.com/booorek/pokedexcli/internal/pokeAPI"
)


func commandExplore(config *config,args[]string) error{
	var locationContent pokeAPI.AreaContent
	var err error

	locationContent, err = config.pokeapiClient.ExploreLocation(&args[0])
	if err != nil {
		return nil
	}
	fmt.Printf("Exploring %s\nFound Pokemon:\n",args[0])
	for _,content := range locationContent.PokemonEncounters{
		fmt.Printf(" - %s\n",content.Pokemon.Name)
	}
	return nil
}

