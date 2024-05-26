package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"lambda-go-series/pkg/handlers"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	lambda.Start(handlers.HelloWorldHandler)
}
