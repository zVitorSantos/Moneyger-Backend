.PHONY: help init build run test clean

help:
    @echo "Usage:"
    @echo "  make init       Initialize project (install dependencies)"
    @echo "  make build      Build the application"
    @echo "  make run        Run the application"
    @echo "  make test       Run tests"
    @echo "  make clean      Clean up"

# Initialize project
init:
    go mod tidy

# Build the application
build:
    go build -o moneyger ./cmd/moneyger
    swag init -g ./cmd/moneyger/main.go -o ./docs/swagger

# Run the application
run:
    go run ./cmd/moneyger/main.go

# Run tests
test:
    go test -v ./...

# Clean up
clean:
    rm -f moneyger
