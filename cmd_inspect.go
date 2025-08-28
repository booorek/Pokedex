package main

import "fmt"



func commandInspect(config *config, args []string) error {
	pokemon, ok := config.coughtPokemons[args[0]]
	if !ok {
		return fmt.Errorf("You have not caught that pokemon")
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n",pokemon.Name,pokemon.Height,pokemon.Weight)
	fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}
