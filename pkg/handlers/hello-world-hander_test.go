package handlers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func TestCanReturnHelloWorldResponse(t *testing.T) {
	response, err := HelloWorldHandler(context.Background(), events.APIGatewayProxyRequest{})

	if err != nil {
		t.Fatalf("Failed to recieve response from handler, err: %v", err)
	}

	if response.StatusCode != 200 {
		t.Fatalf("Expected response code 200,  got %d", response.StatusCode)
	}
}
