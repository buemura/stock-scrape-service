package handlers

import (
	"context"
	"net/http"
	"scraper-service/cmd"
	"scraper-service/infra/http/lambda/helpers"

	"github.com/aws/aws-lambda-go/events"
)

type ScrapeHandlerResponse struct {
	Statue string `json:"statue"`
	Data   any    `json:"data"`
}

func ScrapeHandler(ctx context.Context, request events.APIGatewayProxyRequest) (string, error) {
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

	response := ScrapeHandlerResponse{
		Statue: "success",
		Data:   result,
	}

	return helpers.HandleHttpResponse(http.StatusOK, response)
}
