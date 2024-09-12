package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Explore(area string) ([]Pokemon, error) {
	url := fmt.Sprintf("%v/%v", Urls.Location, area)
	body, ok := ApiCache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		body, err = io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			return nil, errors.New(fmt.Sprintf("Error response from webserver:\nURL: %v\nCode: %v", url, res.StatusCode))
		}
		ApiCache.Add(url, body)
	}

	var areaResponse locationAreaResponse 
	err := json.Unmarshal(body, &areaResponse)
	if err != nil { 
		return nil, err
	}

	var pokemon []Pokemon
	for _, enc := range areaResponse.PokemonEncounters {
		pokemon = append(pokemon, enc.Pokemon)
	}

	return pokemon, nil
}
