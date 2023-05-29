package routesPositions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Laodeus/glt/utils"
	"github.com/Laodeus/glt/utils/tokenUtils"
)

type PositionRequest struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func SendPosition(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {

		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var positionRequest PositionRequest
	err := json.NewDecoder(request.Body).Decode(&positionRequest)
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

	_, err = db.Db.Exec("INSERT INTO location (user_id, lat, lon, time) VALUES ($1, $2, $3, $4)", userid, positionRequest.Lat, positionRequest.Lon, time.Now())

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(err.Error()))
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte("position recorded successfully"))
}
