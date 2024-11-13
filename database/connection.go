package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "pokemon.db")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	log.Println("Conexión a la base de datos exitosa")
}

func Migrate() {
	query := `
    CREATE TABLE IF NOT EXISTS pokemon (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        type TEXT NOT NULL,
        ability TEXT NOT NULL,
        weight INTEGER NOT NULL,
        pokemon_id INTEGER UNIQUE NOT NULL
    );
    `
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Error al migrar la base de datos:", err)
	}
	log.Println("Migración de la base de datos exitosa")
}

func InsertPokemon(name, pType, ability string, weight, pokemonID int) error {
	var exists bool
	queryCheck := `SELECT EXISTS(SELECT 1 FROM pokemon WHERE pokemon_id = ?)`
	err := DB.QueryRow(queryCheck, pokemonID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("el Pokémon con pokemon_id %d ya existe", pokemonID)
	}
	queryInsert := `INSERT INTO pokemon (name, type, ability, weight, pokemon_id) VALUES (?, ?, ?, ?, ?)`
	_, err = DB.Exec(queryInsert, name, pType, ability, weight, pokemonID)
	return err
}

func GetPokemonByID(id string) (map[string]interface{}, error) {
	query := `SELECT id, name, type, ability, weight, pokemon_id FROM pokemon WHERE id = ? OR pokemon_id = ?`
	row := DB.QueryRow(query, id, id)

	var pokemon = make(map[string]interface{})
	var temp struct {
		ID        int
		Name      string
		Type      string
		Ability   string
		Weight    int
		PokemonID int
	}
	err := row.Scan(&temp.ID, &temp.Name, &temp.Type, &temp.Ability, &temp.Weight, &temp.PokemonID)
	if err != nil {
		return nil, err
	}
	pokemon["id"] = temp.ID
	pokemon["name"] = temp.Name
	pokemon["type"] = temp.Type
	pokemon["ability"] = temp.Ability
	pokemon["weight"] = temp.Weight
	pokemon["pokemon_id"] = temp.PokemonID
	return pokemon, nil
}
