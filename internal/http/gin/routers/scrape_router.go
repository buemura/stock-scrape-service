package routers

import (
	"scraper-service/handlers"
	"scraper-service/pkg"

	"github.com/gin-gonic/gin"
)

type ScrapeRequestBody struct {
    Stocks []string `json:"stocks"`
}

type ScrapeRouterResponse struct {
	Statue string `json:"statue"`
	Data   []pkg.StockMarketPrice    `json:"data"`
}

func scrapeRouter(router *gin.RouterGroup) {
	router.GET("/scrape", func(c *gin.Context) {
		data := handlers.ScrapeAuto()
		c.JSON(200, gin.H{
			"status": "success",
			"data": data,
		})
	})

	router.POST("/scrape", func(c *gin.Context) {
		var requestBody ScrapeRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"data": "Failed to get request body",
			})
			return
		}

		data := handlers.ScrapeManual(requestBody.Stocks)
		c.JSON(200, gin.H{
			"status": "success",
			"data": data,
		})
	})
}
