package routesPositions

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Laodeus/glt/utils"
)

type UserMovement struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	Time         time.Time `json:"time"`
	Lat          float64   `json:"lat"`
	Lon          float64   `json:"lon"`
	VehiculeId   *int      `json:"vehicule_id,omitempty"`
	VehiculeName *string   `json:"vehicule_name,omitempty"`
	VehiculeType *string   `json:"vehicule_type,omitempty"`
}

type MouvementsRequest struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func GetUserMovements(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	segments := strings.Split(request.URL.Path, "/")
	if len(segments) < 4 || segments[len(segments)-1] != "movements" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(segments[len(segments)-2])
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	var mouvementsRequest MouvementsRequest
	err = json.NewDecoder(request.Body).Decode(&mouvementsRequest)

	db, err := utils.GetDb()
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Db.Query(`SELECT DISTINCT l.id, l.user_id, l.time, l.lat, l.lon, v.id, v.name, v.type 
	FROM location l
	LEFT JOIN vehicules_usage vu ON l.user_id = vu.user_id AND l.time >= vu.time AND vu.usage = 'take'
	LEFT JOIN vehicules v ON vu.vehicules_id = v.id
	WHERE l.user_id = $1 AND l.time BETWEEN $2 AND $3`, userId, mouvementsRequest.Start, mouvementsRequest.End)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	movements := make([]UserMovement, 0)
	for rows.Next() {
		var movement UserMovement
		err := rows.Scan(&movement.Id, &movement.UserId, &movement.Time, &movement.Lat, &movement.Lon, &movement.VehiculeId, &movement.VehiculeName, &movement.VehiculeType)
		if err != nil && err != sql.ErrNoRows {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			return
		}

		movements = append(movements, movement)
	}

	if err = rows.Err(); err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(responseWriter).Encode(movements)
}
