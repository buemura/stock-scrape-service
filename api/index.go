package handler

import (
	"net/http"
	"scraper-service/internal/http/utils"
)

type RootResponse struct {
	Status string `json:"status"`
	Data any `json:"data"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	utils.HandleSuccessResponse(w, http.StatusOK, "hello from api")
	return
}
