package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SuccessResponse struct {
	Status string `json:"status"`
	Data any `json:"data"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func HandleSuccessResponse(w http.ResponseWriter, status int, data any) {
	res, _ := json.Marshal(&SuccessResponse{
		Status: "success",
		Data: data,
	})

	w.WriteHeader(status)
	fmt.Fprint(w, string(res))
}

func HandleErrorResponse(w http.ResponseWriter, status int, message string) {
	res, _ := json.Marshal(&ErrorResponse{
		Status: "error",
		Message: message,
	})

	w.WriteHeader(status)
	fmt.Fprint(w, string(res))
}
