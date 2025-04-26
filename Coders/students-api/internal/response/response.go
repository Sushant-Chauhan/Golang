package response

import (
	"encoding/json"
	"net/http"
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
