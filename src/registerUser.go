package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func RegisterUser(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user User
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

	db, err := getDb()
	if err != nil {
		fmt.Println("Error creating database connection :", err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.db.Exec("INSERT INTO users (login, password) VALUES ($1, $2)", user.Login, hashedPassword)
	if err != nil {
		fmt.Println("Error inserting user datas :", err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte("User recorded"))
}
