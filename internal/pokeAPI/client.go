package pokeAPI

import (
	"net/http"
	"time"

	"github.com/booorek/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, cacheTimeout time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheTimeout),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
