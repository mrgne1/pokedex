package pokeapi

type MapStepper struct {
	Count          int
	Next           *string
	Prev           *string
	Current        []Location
	VisiblePokemon []PokemonLocation
}

type Location struct {
	Name string
	Url  string
}

type allLocationAreaResponse struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type PokemonLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type pokemonEncounter struct {
	Pokemon PokemonLocation `json:"pokemon"`
}

type locationAreaResponse struct {
	Id                int                `json:"id"`
	Name              string             `json:"name"`
	GameIndex         int                `json:"game_index"`
	Location          Location           `json:"location"`
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

type Ability struct {
	IsHidden bool     `json:"is_hidden"`
	Slot     int      `json:"slot"`
	Ability  Location `json:"ability"`
}

type PokemonStat struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	Stat     Location `json:"stat"`
}

type PokemonType struct {
	Slot int      `json:"slot"`
	Type Location `json:"type"`
}

type Pokemon struct {
	Id             int           `json:"id"`
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	IsDefault      bool          `json:"is_default"`
	Order          int           `json:"order"`
	Weight         int           `json:"weight"`
	Abilities      []Ability     `json:"abilities"`
	Forms          []Location    `json:"forms"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
	// TODO: Finish adding pokemon
}
