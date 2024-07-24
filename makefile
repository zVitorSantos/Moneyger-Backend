.PHONY: all build run clean swag test docker-up docker-down

all: run

build: swag
	@go build -o moneyger cmd/moneyger/main.go

run: swag
	@go run cmd/moneyger/main.go

clean:
	@rm -f moneyger

swag:
	@swag init -g cmd/moneyger/main.go --output docs

test:
	@go test -v ./...

docker-up:
	@cd deployments && docker-compose up --build

docker-down:
	@cd deployments && docker-compose down
