package main

import (
	"context"
	"net/http"

	"scraper-service/infra/http/lambda/handlers"
	"scraper-service/infra/http/lambda/helpers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func router(ctx context.Context, request events.APIGatewayProxyRequest) (string, error) {
    switch request.Path {
    case "/health":
        return handlers.HealtHandler(ctx, request)
    case "/scrape":
        return handlers.ScrapeHandler(ctx, request)
    default:
        return helpers.HandleHttpResponse(http.StatusNotFound, "Route not found")
    }
}

func main() {
    lambda.Start(router)
}