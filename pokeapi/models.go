package pokeapi

type MapStepper struct {
	Count int
	Next  *string
	Prev  *string
}

type LocationArea struct {
	Name string
	Url  string
}

type allLocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type pokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type locationAreaResponse struct {
	Id                int                `json:"id"`
	Name              string             `json:"name"`
	GameIndex         int                `json:"game_index"`
	Location          LocationArea       `json:"location"`
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}
