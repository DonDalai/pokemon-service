package routes

import (
	"net/http"
	"pokemon-service/database"
	"pokemon-service/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/pokemon/:id", func(c *gin.Context) {
		id := c.Param("id")
		pokemon, err := database.GetPokemonByID(id) // Llamamos directamente al repositorio
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Pokémon no encontrado"})
			return
		}
		c.JSON(http.StatusOK, pokemon)
	})

	router.POST("/pokemon", func(c *gin.Context) {
		var input struct {
			PokemonID int `json:"pokemon_id" binding:"required"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Llamada directa al servicio para obtener datos de la PokeAPI
		pokemonData, err := services.FetchPokemonData(input.PokemonID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener datos de la PokeAPI"})
			return
		}

		// Insertamos el Pokémon directamente usando el repositorio
		err = database.InsertPokemon(pokemonData.Name, pokemonData.Types[0].Type.Name, pokemonData.Abilities[0].Ability.Name, pokemonData.Weight, input.PokemonID)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "El Pokémon ya existe"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Pokémon insertado con éxito", "data": pokemonData})
	})

	return router
}
