package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonAPIResponse struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Types  []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
}

func FetchPokemonData(pokemonID int) (*PokemonAPIResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", pokemonID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al consumir la PokeAPI: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("la PokeAPI devolvió el estado: %s", resp.Status)
	}

	var result PokemonAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta de la PokeAPI: %w", err)
	}

	return &result, nil
}
