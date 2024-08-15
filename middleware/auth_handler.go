package middleware

import (
	"chat-app/helper"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func AuthHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(helper.SigningKey), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "unable to parse the token")
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "unable to parse the token")
			return
		}

		next(w, r)
	}
}
