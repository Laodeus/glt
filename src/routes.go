package main

import (
	"net/http"
)

func Routes() *http.ServeMux {

	router := http.NewServeMux()

	// Route POST /api/v1/users/register
	router.HandleFunc("/api/v1/users/register", RegisterUser)

	return router
}
