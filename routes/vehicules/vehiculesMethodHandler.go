package routesVehicules

import (
	"net/http"
)

func VehiculesMethodHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		PostNewVehicules(responseWriter, request)
		return
	} else if request.Method == http.MethodGet {
		GetVehiculesList(responseWriter, request)
	} else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
