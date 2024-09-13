package commands

import "fmt"

func commandPokedex(_... string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range UserState.CaughtPokemon {
		fmt.Printf("  - %v\n", pokemon.Name)
	}
	return nil
}
