package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Bilal-Ahmed4/to-do-app/internal/repository"
	"github.com/Bilal-Ahmed4/to-do-app/internal/response"
	"github.com/jackc/pgx/v5/pgxpool"
)

type todo struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func CreateNewTodo(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty response")))
			return
		}
		//now i have to create the
		models, err := repository.CreateTodo(pool, todo.Title, todo.Completed)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, models)
	}
}

func GetTodos(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := repository.GetTodos(pool)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, fmt.Errorf("Unable to get todos %s", err))
			return
		}

		response.WriteJson(w, http.StatusFound, todos)

	}
}

func GetTodosById(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Unable to convert the id of todo %s", err)))
			return
		}

		todo, err := repository.GetTodoById(pool, id)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, fmt.Errorf("Unable to get the todo %s", err))
			return
		}

		response.WriteJson(w, http.StatusFound, todo)
	}
}
