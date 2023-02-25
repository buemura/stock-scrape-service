package handlers

import (
	"scraper-service/pkg"

	"github.com/gin-gonic/gin"
)

type ScrapeHandlerResponse struct {
	Statue string `json:"statue"`
	Data   any    `json:"data"`
}

func ScrapeHandler(router *gin.Engine) (gin.IRoutes) {
	return router.GET("/scrape", func(c *gin.Context) {
		stocks := pkg.GetStocksTickers()
		result := pkg.ScrapeStocks(stocks)

		c.JSON(200, gin.H{
			"status": "success",
			"data": result,
		})
	})
}
