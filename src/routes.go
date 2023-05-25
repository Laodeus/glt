package main

import (
	"net/http"
)

func Routes() *http.ServeMux {

	router := http.NewServeMux()

	return router
}
