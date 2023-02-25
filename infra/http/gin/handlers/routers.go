package handlers

import "github.com/gin-gonic/gin"

func SetupRouters(router *gin.Engine) {
	HealthHandler(router)
	ScrapeHandler(router)
}