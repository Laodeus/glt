package routesVehicules

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Laodeus/glt/routes/vehiculesUsage"
	"github.com/Laodeus/glt/utils"
	"github.com/Laodeus/glt/utils/tokenUtils"
)

type VehiculeRequest struct {
	Id int `json:"id"`
}

func TakeVehicule(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var vehiculeRequest VehiculeRequest
	err := json.NewDecoder(request.Body).Decode(&vehiculeRequest)
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

	var dbVehicules DbVehicules
	err = db.Db.QueryRow("SELECT * FROM vehicules WHERE id = $1 limit 1", vehiculeRequest.Id).Scan(&dbVehicules.Id, &dbVehicules.Name, &dbVehicules.Type)
	if err != nil {
		responseWriter.WriteHeader(http.StatusConflict)
		responseWriter.Write([]byte(err.Error()))
		return
	} else if dbVehicules.Id == 0 {
		responseWriter.WriteHeader(http.StatusNotFound)
		responseWriter.Write([]byte("Unknow vehicule id."))
	}

	reqToken := request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		http.Error(responseWriter, "Invalid token", http.StatusUnauthorized)
		return
	}

	reqToken = strings.TrimSpace(splitToken[1])
	userid, err := tokenUtils.ParseToken(reqToken)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(err.Error()))
		return
	}

	var usage vehiculesUsage.VehiculesUsageDb
	err = db.Db.QueryRow("SELECT * FROM vehicules_usage WHERE user_id = $1 ORDER BY time DESC LIMIT 1", userid).Scan(&usage.Id, &usage.UserId, &usage.VehiculesId, &usage.Usage, &usage.Time)
	if err != nil && err != sql.ErrNoRows {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(err.Error()))
		return
	} else if usage.Usage == "take" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte(fmt.Sprintf("you are already in the vehicule %d", usage.Id)))
		return
	}

	_, err = db.Db.Exec("INSERT INTO vehicules_usage (user_id, vehicules_id, usage, time) VALUES ($1, $2, $3, $4)", userid, dbVehicules.Id, "take", time.Now())

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(err.Error()))
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte(fmt.Sprintf("you take the vehicule %d, %s of type %s", dbVehicules.Id, dbVehicules.Name, dbVehicules.Type)))
}
