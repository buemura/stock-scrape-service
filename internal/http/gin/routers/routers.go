package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouters(router *gin.Engine) {
	HealthRouter(router)
	ScrapeRouter(router)
}