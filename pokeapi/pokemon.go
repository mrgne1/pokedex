package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPokemon(name string) (Pokemon, error) {
	url := fmt.Sprintf("%v/%v", Urls.Pokemon, name)
	body, ok := ApiCache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()

		ApiCache.Add(url, body)
	}

	var p Pokemon
	err := json.Unmarshal(body, &p)
	if err != nil {
		return Pokemon{}, err
	}
	
	return p, nil
}
