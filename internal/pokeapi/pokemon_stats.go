package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPokemon -
func (c *Client) GetPokemon(pokemonName string) (PokemonStats, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok { //Use this if cached
		pokemonResp := PokemonStats{} //Still need to unmarshal
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return PokemonStats{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonStats{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonStats{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonStats{}, err
	}

	pokemonResp := PokemonStats{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonStats{}, err
	}

	c.cache.Add(url, dat) //Add new value to cache
	return pokemonResp, nil
}
