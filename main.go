package main

import (
	"time"

	"github.com/oday/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeapiClient,
	}

	startRepl(cfg)
}
