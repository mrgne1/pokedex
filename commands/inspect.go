package commands

import (
	"fmt"
	"pokedex/pokeapi"
	"slices"
)

func commandInspect(args ...string) error {
	name := args[0]

	idx := slices.IndexFunc(UserState.CaughtPokemon, func(p pokeapi.Pokemon) bool { return p.Name == name })
	if idx < 0 {
		fmt.Println("you have not caught that pokemon")
	} else {
		pokemon := UserState.CaughtPokemon[idx]
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, typ := range pokemon.Types {
			fmt.Printf("  - %v\n", typ.Type.Name)
		}
	}
		// TODO: print Types

	return nil
}
