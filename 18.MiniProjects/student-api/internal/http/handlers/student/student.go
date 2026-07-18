package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Bilal-Ahmed4/student-api/internal/storage"
	"github.com/Bilal-Ahmed4/student-api/internal/types"
	"github.com/Bilal-Ahmed4/student-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//we need some kind of json to read the data from the body
		var student types.Student

		//now we have to use the decoder to read the
		// NewDecoder accepts io Reader and the decode accept var and return error
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty response")))
			return
		}

		// request validation
		if err := validator.New().Struct(student); err != nil {
			ValidateErr := err.(validator.ValidationErrors)
			fmt.Println(student)
			response.WriteJson(w, http.StatusBadRequest, response.ValidatorError(ValidateErr))
			return
		}

		slog.Info("creating the student")
		lastid, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		slog.Info("user created successfully", slog.String("userId", fmt.Sprint(lastid)))
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}
		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastid})

	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		//now we have to fetch the student by id
		student, err := storage.GetStudentById(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w,http.StatusOK,student)
		
	}
}
