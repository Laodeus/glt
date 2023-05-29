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
		login text UNIQUE,
		password text
	);

	CREATE TABLE IF NOT EXISTS vehicules (
		id SERIAL PRIMARY KEY,
		name text UNIQUE,
		type text
	);

	CREATE TABLE IF NOT EXISTS location (
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		time TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	
	CREATE TYPE usage_enum AS ENUM ('take', 'leave');

	CREATE TABLE IF NOT EXISTS vehicules_usage (
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		vehicules_id INTEGER,
		usage usage_enum,
		time TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (vehicules_id) REFERENCES vehicules(id)
	);

	ALTER TABLE "public"."location" ADD COLUMN "lat" double precision;
	ALTER TABLE "public"."location" ADD COLUMN "lon" double precision;
		
		`)
	if err != nil {
		return err
	}

	// Autres opérations de séquençage...

	return nil
}
