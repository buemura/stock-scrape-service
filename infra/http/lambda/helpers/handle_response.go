package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HandleHttpResponse(statusCode int, response any) (string, error) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	
	data := &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(jsonResponse),
	}
	return fmt.Sprintf("%v", data), nil
}