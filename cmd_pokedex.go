package main

import "fmt"



func commandPokedex(config *config, args []string) error {
	if len(config.coughtPokemons)==0{
		fmt.Printf("Your Pokedex is empty!\n")
		return nil
	}
	fmt.Printf("Your Pokedex:\n")
	for _,pokemon := range config.coughtPokemons{
		fmt.Printf(" - %s",pokemon.Name)
	}
	fmt.Printf("\n")
	return nil
}
