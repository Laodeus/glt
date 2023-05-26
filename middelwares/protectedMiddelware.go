package middlewares

import (
	"fmt"
	"net/http"

	"strings"

	"github.com/Laodeus/glt/utils/tokenUtils"
	"github.com/golang-jwt/jwt"
)

// handle middelware
func ProtectedMiddelware(next http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {

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

		token, err := tokenUtils.ParseToken(reqToken)

		if err != nil {
			http.Error(responseWriter, "Invalid token", http.StatusUnauthorized)
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
