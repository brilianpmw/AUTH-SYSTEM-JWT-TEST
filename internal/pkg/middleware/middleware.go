package middleware

import (
	"errors"
	"net/http"

	"github.com/brilianpmw/linknau/presentation"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string) (*presentation.Claims, error) {
	claims := &presentation.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return presentation.JwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("unauthorized")
	}

	return claims, nil
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		_, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	})
}
