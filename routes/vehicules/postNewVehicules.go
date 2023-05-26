package routesVehicules

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Laodeus/glt/utils"
)

func PostNewVehicules(responseWriter http.ResponseWriter, request *http.Request) {
	var vehicule Vehicules
	err := json.NewDecoder(request.Body).Decode(&vehicule)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("bad body"))
		return
	}

	db, err := utils.GetDb()
	if err != nil {
		fmt.Println("Error creating database connection :", err)
		return
	}

	defer db.Close()

	_, err = db.Db.Exec("INSERT INTO  vehicules (name, type) VALUES ($1, $2)", vehicule.Name, vehicule.Type)
	if err != nil {
		responseWriter.WriteHeader(http.StatusConflict)
		responseWriter.Write([]byte(err.Error()))
		return
	} else {
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write([]byte("Vehicule added"))
	}
}
