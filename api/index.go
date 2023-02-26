package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RootResponse struct {
	Status string `json:"status"`
	Data any `json:"data"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	response := RootResponse{
		Data: "hello",
		Status: "success",
	}
	jsonResponse, _ := json.Marshal(response)
	fmt.Fprintf(w, string(jsonResponse))
}
