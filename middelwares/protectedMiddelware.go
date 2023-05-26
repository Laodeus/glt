package middlewares

import (
	"net/http"
	"strings"

	"github.com/Laodeus/glt/utils/tokenUtils"
)

// handle middelware
func ProtectedMiddelware(next http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {

		reqToken := request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			http.Error(responseWriter, "Invalid token", http.StatusUnauthorized)
			return
		}

		reqToken = strings.TrimSpace(splitToken[1])

		_, err := tokenUtils.ParseToken(reqToken)

		if err != nil {
			http.Error(responseWriter, "Invalid token", http.StatusUnauthorized)
			return
		}

		next(responseWriter, request)
	}
}
