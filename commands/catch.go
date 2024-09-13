package commands

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"pokedex/pokeapi"
)

func commandCatch(args ...string) error {

	if len(args) != 1 {
		return errors.New("Must give pokemon name")
	}
	name := args[0]

	if !canSee(name) {
		fmt.Printf("Can't see a %v, try exploring another location\n", name)
		return nil
	}

	pokemon, err := pokeapi.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v\n", pokemon.Name)

	if rand.Intn(100) > pokeProb(pokemon) {
		fmt.Printf("%v was caught\n", pokemon.Name)
		UserState.CaughtPokeon = append(UserState.CaughtPokeon, pokemon)
	} else {
		fmt.Printf("%v escaped\n", pokemon.Name)
	}

	return nil
}

func pokeProb(p pokeapi.Pokemon) int {
	return int(100.0 * math.Tanh(float64(p.BaseExperience)/ 100.0))
}

func canSee(name string) bool {
	for _, p := range mapStepper.VisiblePokemon {
		if name == p.Name {
			return true
		}
	}
	return false
}
