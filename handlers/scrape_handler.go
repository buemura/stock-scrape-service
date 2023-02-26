package handlers

import "scraper-service/pkg"

func ScrapeAuto() []pkg.StockMarketPrice {
	stocks := pkg.GetStocksTickers()
	result := pkg.ScrapeStocks(stocks)
	return result
}

func ScrapeManual(stocks []string) []pkg.StockMarketPrice {
	result := pkg.ScrapeStocks(stocks)
	return result
}