package routesVehicules

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Laodeus/glt/utils"
)

func GetVehiculesList(responseWriter http.ResponseWriter, request *http.Request) {

	db, err := utils.GetDb()
	if err != nil {
		fmt.Println("Error creating database connection :", err)
		return
	}

	defer db.Close()

	var vehicules []DbVehicules
	rows, err := db.Db.Query("SELECT * from vehicules ")
	for rows.Next() {
		var vehic DbVehicules
		err := rows.Scan(&vehic.Id, &vehic.Name, &vehic.Type)
		if err != nil {

		}
		vehicules = append(vehicules, vehic)
	}

	vehiculesJSON, err := json.Marshal(vehicules)
	if err != nil {
		// handle error here
		fmt.Println("Error converting to JSON:", err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err != nil {
		responseWriter.WriteHeader(http.StatusConflict)
		responseWriter.Write([]byte(err.Error()))
		return
	} else {
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write(vehiculesJSON)
	}
}
