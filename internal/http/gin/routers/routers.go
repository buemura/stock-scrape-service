package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrRouter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"message": "Route not found",
	})
}

func SetupRouters(router *gin.RouterGroup) {
	healthRouter(router)
	scrapeRouter(router)
}
