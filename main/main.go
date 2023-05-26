package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Laodeus/glt/routes"
	"github.com/Laodeus/glt/utils"
)

func main() {
	// load the env variables
	utils.LoadEnv()

	port := os.Getenv("SERVER_PORT")
	log.Println(port)

	// connect to db
	db, err := utils.GetDb()
	if err != nil {
		fmt.Println("Error creating database connection :", err)
		return
	}
	defer db.Close()

	router := routes.Routes()

	fmt.Println(fmt.Sprintf("Server started at port %s", port))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
