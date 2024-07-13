.PHONY: all build run clean swag test

all: run

build: swag build
	@go build -o moneyger cmd/moneyger/main.go

run:swag
	@go run cmd/moneyger/main.go

clean:
	@rm -f moneyger

swag:
	@swag init -g cmd/moneyger/main.go --output docs

test:
	@go test -v ./...
