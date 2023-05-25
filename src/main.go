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

	// definition des routes
	router := Routes()

	// Message indiquant que le serveur a démarré
	fmt.Println(fmt.Sprintf("Le serveur a démarré sur le port %s", port))

	// Démarrage du serveur
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
