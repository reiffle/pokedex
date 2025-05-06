package main

import (
	"time"

	"github.com/reiffle/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.PokemonStats{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
