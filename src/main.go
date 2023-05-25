package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// load the env variables
	LoadEnv()

	port := os.Getenv("SERVER_PORT")
	log.Println(port)

	// connect to db
	db, err := getDb()
	if err != nil {
		fmt.Println("Error creating database connection :", err)
		return
	}
	defer db.Close()

	router := Routes()

	fmt.Println(fmt.Sprintf("Server started at port %s", port))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
