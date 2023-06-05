package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {
	locationAreas, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas:")
	for _, area := range locationAreas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = locationAreas.Next
	cfg.prevLocationAreaURL = locationAreas.Previous
	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	locationAreas, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas:")
	for _, area := range locationAreas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = locationAreas.Next
	cfg.prevLocationAreaURL = locationAreas.Previous
	return nil
}
