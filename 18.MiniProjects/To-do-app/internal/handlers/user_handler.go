package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Bilal-Ahmed4/to-do-app/internal/config"
	"github.com/Bilal-Ahmed4/to-do-app/internal/models"
	"github.com/Bilal-Ahmed4/to-do-app/internal/repository"
	"github.com/Bilal-Ahmed4/to-do-app/internal/response"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationRequest struct {
	Email    string
	Password string
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func CreateUser(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user RegistrationRequest
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Unable to decode the body%s", err)))
			return
		}

		if len(user.Password) < 6 {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Password must be at least 6 chars")))
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

	}
}
