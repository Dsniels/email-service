package pkg

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Data T
}

type ErrorResponse struct {
	Status  int
	Message string
}

func WriteReponse[T any](w http.ResponseWriter, status int, data T) {
	res := &Response[T]{
		Data: data,
	}
	jsn, _ := json.Marshal(res)
	w.WriteHeader(status)
	w.Write(jsn)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	res := &ErrorResponse{
		Status:  status,
		Message: message,
	}

	jsn, _ := json.Marshal(res)
	w.WriteHeader(status)
	w.Write(jsn)
}
