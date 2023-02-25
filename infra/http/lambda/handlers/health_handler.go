package handlers

import (
	"context"
	"net/http"
	"scraper-service/infra/http/lambda/helpers"

	"github.com/aws/aws-lambda-go/events"
)

type HealtHandlerResponse struct {
	Message string `json:"message"`
}

func HealtHandler(ctx context.Context, request events.APIGatewayProxyRequest) (string, error) {
	response := HealtHandlerResponse{
		Message: "API working",
	}
	return helpers.HandleHttpResponse(http.StatusOK, response)
}