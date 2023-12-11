package middleware

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(jwtSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Get user information from token
			//claims := token.Claims.(jwt.MapClaims)
			//userID := claims["id"].(string)

			// Retrieve user from database or cache
			// user, err := nil, nil

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(r.Context(), "user", "")
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
