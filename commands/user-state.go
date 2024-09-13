package commands

import "pokedex/pokeapi"
 

type Pokedex struct {
	CaughtPokemon []pokeapi.Pokemon
}

var UserState = Pokedex {
	CaughtPokemon: make([]pokeapi.Pokemon, 10),
}
