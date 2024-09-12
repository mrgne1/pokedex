package commands

import (
	"errors"
	"fmt"
	"pokedex/pokeapi"
)

func commandExplore(args ...string) error {
	if len(args) == 0 {
		return errors.New("Must provide area")
	}
	if len(args) > 1 {
		return errors.New("Can only explore one area")
	}

	area := args[0]
	fmt.Printf("Exploring %v\n", area)

	pokemon, err := pokeapi.Explore(area)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, poke := range pokemon {
		fmt.Printf(" - %v\n", poke.Name)
	}
	return nil
}
