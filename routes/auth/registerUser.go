package routesAuth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Laodeus/glt/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user DbLogin
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("bad body"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	db, err := utils.GetDb()
	if err != nil {
		fmt.Println("Error creating database connection :", err)
		return
	}
	defer db.Close()

	_, err = db.Db.Exec("INSERT INTO users (login, password) VALUES ($1, $2)", user.Login, hashedPassword)
	if err != nil {
		fmt.Println("Error inserting user datas :", err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte("User recorded"))
}
