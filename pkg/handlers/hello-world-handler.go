package handlers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

func HelloWorldHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Perform some work here

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "Hello world",
	}

	return response, nil
}
