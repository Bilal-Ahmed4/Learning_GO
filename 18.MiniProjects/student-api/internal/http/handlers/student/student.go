package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Bilal-Ahmed4/student-api/internal/types"
	"github.com/Bilal-Ahmed4/student-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	})
}
