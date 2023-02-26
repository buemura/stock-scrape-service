package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"scraper-service/handlers"
)

type HealthResponse struct {
	Status string `json:"status"`
	Data any `json:"data"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	data := handlers.HealthHandler()
	response := HealthResponse{
		Status: "success",
		Data:  data,
	}
	
	jsonResponse, _ := json.Marshal(response)
	fmt.Fprintf(w, string(jsonResponse))
}