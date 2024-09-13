package commands

import "pokedex/pokeapi"
 

type Pokedex struct {
	CaughtPokeon []pokeapi.Pokemon
}

var UserState = Pokedex {
	CaughtPokeon: make([]pokeapi.Pokemon, 10),
}
