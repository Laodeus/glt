package routesAuth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Laodeus/glt/utils"
	"github.com/Laodeus/glt/utils/tokenUtils"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(responseWriter http.ResponseWriter, request *http.Request) {
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

	db, err := utils.GetDb()
	if err != nil {
		fmt.Println("Error creating database connection :", err)
		return
	}
	defer db.Close()
	var retrievedUser DbLoginRow
	err = db.Db.QueryRow("SELECT * FROM users WHERE login = $1 limit 1", user.Login).Scan(&retrievedUser.Id, &retrievedUser.Login, &retrievedUser.Password)

	isUserPasswordOk := bcrypt.CompareHashAndPassword([]byte(retrievedUser.Password), []byte(user.Password))

	if isUserPasswordOk != nil {
		responseWriter.WriteHeader(http.StatusUnauthorized)
		responseWriter.Write([]byte("bad credential"))
	}

	generatedToken, err := tokenUtils.GenerateToken(retrievedUser.Id)

	responseWriter.WriteHeader(http.StatusUnauthorized)
	responseWriter.Write([]byte(fmt.Sprintf("bearer %s", generatedToken)))

}
