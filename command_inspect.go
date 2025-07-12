package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no Pokemon provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("Pokemon have not been caught yet")
	}

	fmt.Printf("Name: %s", pokemon.Name)
	fmt.Printf("Species: %s", pokemon.Species.Name)

	for _, stats := range pokemon.Stats {
		fmt.Printf("- %s\n", stats)
	}

	return nil
}
