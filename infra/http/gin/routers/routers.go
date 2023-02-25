package routers

import (
	"scraper-service/infra/http/gin/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouters(router *gin.Engine) {
	handlers.HealthHandler(router)
	handlers.ScrapeHandler(router)
}