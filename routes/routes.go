package routes

import (
	"net/http"
	"regexp"

	middlewares "github.com/Laodeus/glt/middelwares"
	routesAuth "github.com/Laodeus/glt/routes/auth"
	routeMouvement "github.com/Laodeus/glt/routes/mouvements"
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

	// Route POST and GET /api/v1/vehicles/leave
	router.HandleFunc("/api/v1/vehicles/leave", middlewares.ProtectedMiddelware(routesVehicules.LeaveVehicule))

	// Mouvement
	// Route POST /api/v1/positions
	router.HandleFunc("/api/v1/positions", middlewares.ProtectedMiddelware(routeMouvement.SendPosition))

	//POST /api/v1/users/{userId}/movements
	router.HandleFunc("/api/v1/users/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		match, _ := regexp.MatchString(`^/api/v1/users/[0-9]+/movements$`, path)
		if match {
			middlewares.ProtectedMiddelware(routeMouvement.GetUserMovements)(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	return router
}
