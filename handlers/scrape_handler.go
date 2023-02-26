package handlers

import "scraper-service/pkg"

func ScrapeHandler() []pkg.StockMarketPrice {
	stocks := pkg.GetStocksTickers()
	result := pkg.ScrapeStocks(stocks)
	return result
}