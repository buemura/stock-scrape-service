package routers

import (
	"scraper-service/handlers"
	"scraper-service/pkg"

	"github.com/gin-gonic/gin"
)

type ScrapeRouterResponse struct {
	Statue string `json:"statue"`
	Data   []pkg.StockMarketPrice    `json:"data"`
}

func ScrapeRouter(router *gin.Engine) (gin.IRoutes) {
	return router.GET("/scrape", func(c *gin.Context) {
		data := handlers.ScrapeAuto()
		c.JSON(200, gin.H{
			"status": "success",
			"data": data,
		})
	})
}
