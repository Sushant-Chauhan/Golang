package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Sushant-Chauhan/Go/Coders/students-api/internal/response"
	"github.com/Sushant-Chauhan/Go/Coders/students-api/internal/types"
	"github.com/go-playground/validator"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			// response.WriteJson(w, http.StatusBadRequest, err.Error())
			// "EOF"
			// response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))  //custom json we are returing got Error EOF using GeneralError package
			// {
			//   "status": "ERROR",
			//   "message": "EOF"
			// }
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("body is empty"))) // custom error message as EOF is not user friendly, so we are using 'empty body' as message.
			// {
			// 	"status": "ERROR",
			// 	"message": "empty body"
			// }

			return
		}

		// if there is other error apart from empty body, then we will return that error as it is. No custome message.
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Zero Trust Policy: Validate the data before using it.
		if err := validator.New().Struct(student); err != nil { // this will return error if validation fails. We can use this to check if the data is valid or not.
			validateErrs := err.(validator.ValidationErrors)                                     // this will return the validation errors.
			response.WriteJson(w, http.StatusBadRequest, response.MyValidationErrorFunctionn(validateErrs)) // this will return the validation errors in JSON format.
			return
		}

		// w.Write([]byte("Welcome to Students API !!"))
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
