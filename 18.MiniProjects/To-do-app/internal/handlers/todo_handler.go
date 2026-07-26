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
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type todo struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type UpdateTodo struct {
	Title     *string `json: "title"`
	Completed *bool   `json :"completed"`
}

func CreateNewTodoHandler(pool *pgxpool.Pool) http.HandlerFunc {
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

func GetTodosHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := repository.GetTodos(pool)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, fmt.Errorf("Unable to get todos %s", err))
			return
		}

		response.WriteJson(w, http.StatusFound, todos)

	}
}

func GetTodosByIdHandler(pool *pgxpool.Pool) http.HandlerFunc {
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

func UpdateTodoHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Cannot convert the id from string to int %s", err)))
			return
		}

		//now we will get values from the body
		var updateTodo UpdateTodo
		err = json.NewDecoder(r.Body).Decode(&updateTodo)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("unable to decode the body %s", err)))
		}

		if updateTodo.Title == nil && updateTodo.Completed == nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Provide at least one value to update")))
			return
		}

		existingTodo, err := repository.GetTodoById(pool, id)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("The todo of this id doesnt exists")))
			return
		}

		var title string
		title = existingTodo.Title
		if updateTodo.Title != nil {
			title = *updateTodo.Title
		}

		//now we will update the todos
		var completed bool
		completed = existingTodo.Completed
		if updateTodo.Completed != nil {
			completed = *updateTodo.Completed
		}

		todo, err := repository.UpdateTodo(pool, title, completed, id)
		if err != nil {
			if err == pgx.ErrNoRows {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Error Todo not found%s", err)))
				return
			}
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Unable to get the data from the db %s", err)))
			return
		}

		response.WriteJson(w, http.StatusOK, todo)

	}

}
