package routers

import (
	"net/http"
	"scraper-service/handlers"

	"github.com/gin-gonic/gin"
)

func healthRouter(router *gin.RouterGroup) {
	router.GET("/health", func(c *gin.Context) {
		health := handlers.HealthHandler()
		c.JSON(http.StatusOK, health)
	})
}