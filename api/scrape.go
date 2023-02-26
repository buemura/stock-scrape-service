package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"scraper-service/handlers"
)

type ScrapeResponse struct {
	Status string `json:"status"`
	Data any `json:"data"`
}

func Scrape(w http.ResponseWriter, r *http.Request) {
	data := handlers.ScrapeHandler()
	response := ScrapeResponse{
		Status: "success",
		Data:   data,
	}
	
	jsonResponse, _ := json.Marshal(response)
	fmt.Fprintf(w, string(jsonResponse))
}