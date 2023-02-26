package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Data any `json:"data"`
	Status string `json:"status"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Data: "hello",
		Status: "success",
	}
	jsonResponse, _ := json.Marshal(response)
	fmt.Fprintf(w, string(jsonResponse))
}
