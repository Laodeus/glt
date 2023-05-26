package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	Db *sql.DB
}

func GetDb() (*Database, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success connecting to database")

	return &Database{Db: db}, nil
}

func (d *Database) Close() {
	d.Db.Close()
}
