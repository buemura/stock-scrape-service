package handlers

import (
	"scraper-service/cmd"

	"github.com/gin-gonic/gin"
)

type ScrapeHandlerResponse struct {
	Statue string `json:"statue"`
	Data   any    `json:"data"`
}

func ScrapeHandler(router *gin.Engine) (gin.IRoutes) {
	return router.GET("/scrape", func(c *gin.Context) {
		stocks := []string{
			"VTI",
			"PFF",
			"JEPI",
			"PGX",
			"LQD",
			"BCFF11.SA",
			"BDIF11.SA",
			"BRCR11.SA",
			"BTLG11.SA",
			"HCTR11.SA",
			"HGRE11.SA",
			"IFRA11.SA",
			"KISU11.SA",
			"MXRF11.SA",
			"RZAG11.SA",
			"VISC11.SA",
			"XPCA11.SA",
			"XPLG11.SA",
			"XPML11.SA",
			"BTC-USD",
			"ETH-USD",
			"ADA-USD",
		}
	
		result := cmd.StartScraping(stocks)

		c.JSON(200, gin.H{
			"status": "success",
			"data": result,
		})
	})
}
