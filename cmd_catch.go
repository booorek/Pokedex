package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateChance(baseExperience int, r *rand.Rand) float64 {
	balance := 300.0
	randomVal := 0.5 + r.Float64()*0.5
	return (balance / (float64(baseExperience) + balance)) * randomVal
}

func commandCatch(config *config, args []string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	pokemonInfo, err := config.pokeapiClient.GetPokemonInfo(&args[0])
	if err != nil {
		return fmt.Errorf("Error while acquiring data %s\n", err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	chanceToCatch := generateChance(pokemonInfo.BaseExperience, r)
	roll := r.Float64()

	if chanceToCatch > roll {
		fmt.Printf("Catched\n")
	} else {
		fmt.Printf("Failed while trying to catch %s\n", args[0])
	}

	return nil
}
