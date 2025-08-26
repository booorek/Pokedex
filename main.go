package main

import (
	"time"
	"github.com/booorek/pokedexcli/internal/pokeAPI"
)

func main(){
	client := pokeAPI.NewClient(time.Second*5,time.Minute*5)
	config := &config{
		pokeapiClient: client,
	}
	startPokedex(config)
	
}
