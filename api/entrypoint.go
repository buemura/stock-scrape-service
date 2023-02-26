package api

import (
	"net/http"
	"scraper-service/internal/http/gin/routers"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.New()
	app.NoRoute(routers.ErrRouter)

	router := app.Group("/api")
	routers.SetupRouters(router)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}