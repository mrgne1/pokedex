package pokeapi

import (
	"pokedex/internal/pokecache"
	"time"
)

var ApiCache pokecache.Cache = *pokecache.NewCache(5 * time.Minute)
