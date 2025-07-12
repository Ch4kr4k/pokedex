package main

import (
	"errors"
	"fmt"
	"log"
)

func callBackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Location areas")

	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.revLocationAreaUrl = resp.Previous
	return nil
}

func callBackMapb(cfg *config, args ...string) error {
	if cfg.revLocationAreaUrl == nil {
		return errors.New("First Page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.revLocationAreaUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas")

	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.revLocationAreaUrl = resp.Previous
	return nil
}

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemen in %s\n", locationArea.Names)

	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
