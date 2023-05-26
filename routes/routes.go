package routes

import (
	"net/http"

	routesAuth "github.com/Laodeus/glt/routes/auth"
)

func Routes() *http.ServeMux {

	router := http.NewServeMux()

	// Route POST /api/v1/users/register
	router.HandleFunc("/api/v1/users/register", routesAuth.RegisterUser)

	return router
}
