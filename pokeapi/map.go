package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func NewMapStepper() *MapStepper {
	c := MapStepper{
		Count:          0,
		Next:           &Urls.Location,
		Prev:           &Urls.Location,
		Current:        make([]Location, 10),
		VisiblePokemon: make([]PokemonLocation, 10),
	}
	return &c
}

func (s *MapStepper) GetMap() ([]Location, error) {
	if s.Next != nil {
		return s.getLocationAreas(*s.Next)
	} else {
		return s.getLocationAreas(Urls.Location)
	}
}

func (s *MapStepper) GetMapb() ([]Location, error) {
	if s.Prev != nil {
		return s.getLocationAreas(*s.Prev)
	} else {
		return s.getLocationAreas(Urls.Location)
	}
}

func (s *MapStepper) getLocationAreas(url string) ([]Location, error) {
	body, ok := ApiCache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return nil, errors.New(fmt.Sprintf("Error response from webserver:\nURL: %v\nCode: %v", url, res.StatusCode))
		}

		if err != nil {
			return nil, err
		}
		ApiCache.Add(url, body)
	}

	var areas allLocationAreaResponse
	json.Unmarshal(body, &areas)
	s.Count = areas.Count
	s.Next = areas.Next
	s.Prev = areas.Previous
	s.Current = areas.Results

	return areas.Results, nil
}

func (s *MapStepper) Print() {
	var next, prev string
	if s.Next != nil {
		next = *s.Next
	} else {
		next = "nil"
	}
	if s.Prev != nil {
		prev = *s.Prev
	} else {
		prev = "nil"
	}
	fmt.Printf("{\n  Count: %v\n  Prev: %v\n  Next: %v\n}\n", s.Count, next, prev)
}
