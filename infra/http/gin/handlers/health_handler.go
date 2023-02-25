package handlers

import (
	"github.com/gin-gonic/gin"
)

func HealthHandler(router *gin.Engine) (gin.IRoutes) {
	return router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API working",
		})
	})
}