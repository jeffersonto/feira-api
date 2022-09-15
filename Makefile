lint:
	@go mod tidy
	@golangci-lint run --fix
	@golangci-lint run ./...

.PHONY: test
