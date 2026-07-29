package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Bilal-Ahmed4/to-do-app/internal/config"
	"github.com/Bilal-Ahmed4/to-do-app/internal/middleware"
	"github.com/Bilal-Ahmed4/to-do-app/internal/models"
	"github.com/Bilal-Ahmed4/to-do-app/internal/repository"
	"github.com/Bilal-Ahmed4/to-do-app/internal/response"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func CreateUser(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var validate = validator.New()
		var user RegistrationRequest
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Unable to decode the body%s", err)))
			return
		}

		if err := validate.Struct(user); err != nil {
			http.Error(w, "validation failed: "+err.Error(), http.StatusBadRequest)
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Unable to decrypt the password%s", err)))
			return
		}

		userModel := &models.User{
			Email:    user.Email,
			Password: string(hashedPassword),
		}

		createdUser, err := repository.CreateUser(pool, userModel)

		if err != nil {
			if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
				response.WriteJson(w, http.StatusBadRequest, fmt.Errorf("Email already registred"))
				return
			}

			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(fmt.Errorf("Unable to create the user")))
			return
		}

		response.WriteJson(w, http.StatusCreated, createdUser)

	}
}

func LoginHandler(pool *pgxpool.Pool, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var validate = validator.New()
		var req LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Unable to decode the boyd %s", err)))
			return
		}
		//now we validate
		if err := validate.Struct(req); err != nil {
			http.Error(w, "validation failed: "+err.Error(), http.StatusBadRequest)
			return
		}

		//now we get the email from the db and match the password
		user, err := repository.GetUserByEmail(pool, req.Email)
		if err != nil {
			response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(fmt.Errorf("Email or password is wrong %s", err)))
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

		if err != nil {
			response.WriteJson(w, http.StatusUnauthorized, err.Error())
			return
		}

		//now we have confirm the password now we will generate the jwt token
		claims := jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
		}

		//now we will generate the token object for the token string
		tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := tokenObj.SignedString([]byte(cfg.JWTSecret))
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusOK, tokenString)
	}
}

// Test handler to check whether the middleware is working or not
func TestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := r.Context().Value(middleware.UserIDKey).(string)
		if !ok {
			response.WriteJson(w, http.StatusBadRequest, fmt.Errorf("invalid user"))
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]interface{}{
			"message": "Hi how you are doing",
			"user_id": userId,
		})
	}

}
