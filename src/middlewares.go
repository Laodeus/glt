package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// handle middelware
func ProtectedMiddelware(next http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {

		var jwtKey = []byte(os.Getenv("SECRET"))

		// get the token
		reqToken := request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")

		if len(splitToken) != 2 {
			http.Error(responseWriter, "Invalid token", http.StatusUnauthorized)
			return
		}

		reqToken = splitToken[1]
		responseWriter.WriteHeader(http.StatusUnauthorized)
		responseWriter.Write([]byte("not found!"))

		// Parse the token
		token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
			// Make sure the token method conform to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			http.Error(responseWriter, "Invalid token", http.StatusUnauthorized)
			return
		}

		// If the token is valid, pass the request to the next handler
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// You can access the token claims here. For example:
			fmt.Println(claims["foo"], claims["nbf"])
			next(responseWriter, request)
		} else {
			http.Error(responseWriter, "Invalid token", http.StatusUnauthorized)
		}

		// next(responseWriter, request)
	}
}
