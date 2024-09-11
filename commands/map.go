package commands

import (
	"fmt"
	"pokedex/pokeapi"
)

var mapStepper pokeapi.MapStepper = *pokeapi.NewMapStepper()

func commandMap() error {
	locations, err := mapStepper.GetMap()
	if err != nil {
		return err
	}

	for _, loc := range locations {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb() error {
	locations, err := mapStepper.GetMapb()
	if err != nil {
		return err
	}

	for _, loc := range locations {
		fmt.Println(loc.Name)
	}
	return nil
}
