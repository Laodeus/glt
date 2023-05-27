package routes

import (
	"net/http"

	middlewares "github.com/Laodeus/glt/middelwares"
	routesAuth "github.com/Laodeus/glt/routes/auth"
	routesVehicules "github.com/Laodeus/glt/routes/vehicules"
)

func Routes() *http.ServeMux {

	router := http.NewServeMux()

	// Auth
	// Route POST /api/v1/users/register
	router.HandleFunc("/api/v1/users/register", routesAuth.RegisterUser)

	// Route POST /api/v1/users/register
	router.HandleFunc("/api/v1/users/login", routesAuth.LoginUser)

	// Vehicules
	// Route POST and GET /api/v1/vehicles
	router.HandleFunc("/api/v1/vehicles", middlewares.ProtectedMiddelware(routesVehicules.VehiculesMethodHandler))

	// Route POST and GET /api/v1/vehicles
	router.HandleFunc("/api/v1/vehicles/take", middlewares.ProtectedMiddelware(routesVehicules.TakeVehicule))

	return router
}
