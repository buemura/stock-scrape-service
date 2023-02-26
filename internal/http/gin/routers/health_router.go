package routers

import (
	"net/http"
	"scraper-service/handlers"

	"github.com/gin-gonic/gin"
)

func HealthRouter(router *gin.Engine) (gin.IRoutes) {
	health := handlers.HealthHandler()
	
	return router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, health)
	})
}