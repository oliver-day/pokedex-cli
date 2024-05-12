package main

import (
	"time"

	"github.com/oday/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeapiClient,
	}

	startRepl(cfg)
}
