package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type MapStepper struct {
	Count   int
	Next    *string
	Prev    *string
	Results []LocationArea
}

type LocationArea struct {
	Name string
	Url  string
}

type locationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string         `json:"next"`
	Previous *string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

var startURL string = "https://pokeapi.co/api/v2/location-area"

func (s *MapStepper) GetMap() ([]LocationArea, error) {
	if s.Next != nil {
		return s.getLocationAreas(*s.Next)
	} else {
		return s.getLocationAreas(startURL)
	}
}

func (s *MapStepper) GetMapb() ([]LocationArea, error) {
	if s.Prev != nil {
		return s.getLocationAreas(*s.Prev)
	} else {
		return s.getLocationAreas(startURL)
	}
}

func (s *MapStepper) getLocationAreas(url string) ([]LocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Error response from webserver:\nURL: %v\nCode: %v", url, res.StatusCode))
	}

	if err != nil {
		return nil, err
	}

	var areas locationAreaResponse
	json.Unmarshal(body, &areas)
	s.Count = areas.Count
	s.Next = areas.Next
	s.Prev = areas.Previous

	return areas.Results, nil
}

func (s MapStepper) Print() {
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
