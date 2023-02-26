package handler

import (
	"net/http"
	"scraper-service/handlers"
	"scraper-service/internal/http/utils"
)

type ScrapeRequestBody struct {
    Stocks []string `json:"stocks"`
}

func Scrape(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": {
		data := handlers.ScrapeAuto()
		utils.HandleSuccessResponse(w, http.StatusOK, data)
		return
	}
	case "POST": {
		var requestBody ScrapeRequestBody
		utils.GetRequestBody(w, r, &requestBody)
		data := handlers.ScrapeManual(requestBody.Stocks)
		utils.HandleSuccessResponse(w, http.StatusOK, data)
		return
	}
	default: {
		utils.HandleErrorResponse(w, http.StatusNotFound, "Route not found")
		return
	}
	}
}