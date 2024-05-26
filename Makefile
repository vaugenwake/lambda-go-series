build-hello-world:
	@echo "Building hello-world function"
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o ./artifacts/functions/hello-world cmd/functions/hello-world.go

test:
	@go test ./... -v