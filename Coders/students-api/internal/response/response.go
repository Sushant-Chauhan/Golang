package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

type Response struct {
	Status string `json:"custom_status"` //when this field will get serialized, it will be converted to JSON format like we specify in `json` tag. Other it will return same value as it is in Struct field i.e Status.
	Error  string `json:"custom_message"`
}

const (
	StatusOK    = "OK"
	StatusError = "ERROR"
)

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

// ValidationError is used to return validation errors in a custom format.
// It takes a validator.ValidationErrors as input and returns a Response struct with the errors formatted as a string.
func MyValidationErrorFunctionn(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is required field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
