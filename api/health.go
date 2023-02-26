package handler

import (
	"net/http"
	"scraper-service/handlers"
	"scraper-service/internal/http/utils"
)

type HealthResponse struct {
	Status string `json:"status"`
	Data any `json:"data"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": {
		data := handlers.HealthHandler()
		utils.HandleSuccessResponse(w, http.StatusOK, data)
		return
	}
	default: {
		utils.HandleErrorResponse(w, http.StatusNotFound, "Route not found")
		return
	}
	}
	
}