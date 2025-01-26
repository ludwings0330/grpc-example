.PHONY: all client server lint tidy

client:
	@echo "Running client..."
	@cd client && go run ./client.go

server:
	@echo "Running server..."
	@cd server && go run ./server.go

lint:
	@echo "Running golangci-lint..."
	@cd client && golangci-lint run ./... 
	@cd server && golangci-lint run ./...

tidy:
	@echo "Running go mod tidy on all modules..."
	@cd client && go mod tidy
	@cd server && go mod tidy

all: tidy lint server
