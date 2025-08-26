package main

import (
	"time"
	"github.com/booorek/pokedexcli/internal/pokeAPI"
)

func main(){
	client := pokeAPI.NewClient(time.Second*5)
	startPokedex(&client)
	
}
