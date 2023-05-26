package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Laodeus/glt/utils"
)

func main() {

	utils.LoadEnv()

	db, err := utils.GetDb()
	err = seedDatabase(db.Db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Séquençage terminé avec succès !")
}

func seedDatabase(db *sql.DB) error {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		login VARCHAR(50) UNIQUE,
		password VARCHAR(50)
	);

	CREATE TABLE IF NOT EXISTS vehicules (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) UNIQUE,
		type VARCHAR(50)
	);

	CREATE TABLE IF NOT EXISTS location (
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		time TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS vehicules_usage (
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		vehicules_id INTEGER,
		time TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (vehicules_id) REFERENCES vehicules(id)
	);
		
		`)
	if err != nil {
		return err
	}

	// Autres opérations de séquençage...

	return nil
}
