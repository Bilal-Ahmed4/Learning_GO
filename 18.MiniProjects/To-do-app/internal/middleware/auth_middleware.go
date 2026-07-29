package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Bilal-Ahmed4/to-do-app/internal/config"
	"github.com/Bilal-Ahmed4/to-do-app/internal/response"
	"github.com/golang-jwt/jwt"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func AuthMiddleware(cfg *config.Config, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//now we will get the authHeader to get the token
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(fmt.Errorf("Authorization Header Required")))
			return
		}

		const prefix = "Bearer "

		tokenString := strings.TrimPrefix(authHeader, prefix)

		if tokenString == "" || tokenString == authHeader {
			response.WriteJson(w, http.StatusUnauthorized, "Invalid authorization header format")
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(fmt.Errorf("Invalid or expired token")))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims) //asserting the concrete type stored inside this interface is jwt.MapClaims
		if !ok {
			response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(fmt.Errorf("Invalide token claims")))
			return
		}

		userId, ok := claims["user_id"].(string)
		if !ok {
			response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(fmt.Errorf("Invalide token claims")))
			return
		}

		if exp, ok := claims["exp"].(float64); ok {
			expTime := time.Unix(int64(exp), 0)

			if time.Now().After(expTime) {
				response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(fmt.Errorf("Token has expired")))
				return
			}

		}

		//"Take the user ID I just figured out from the token,
		//tuck it inside the request so anyone downstream can grab it,
		// and then hand control over to the next handler in the chain."

		ctx := context.WithValue(r.Context(), UserIDKey, userId)

		next.ServeHTTP(w, r.WithContext(ctx))

	}
}
