package main

import (
	"time"

	"github.com/Ch4kr4k/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaUrl *string
	revLocationAreaUrl  *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	// testHttpRequest()
	cfg := config{
		pokeapiClient:       pokeapi.NewClient(time.Hour),
		nextLocationAreaUrl: nil,
		revLocationAreaUrl:  nil,
		caughtPokemon:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
