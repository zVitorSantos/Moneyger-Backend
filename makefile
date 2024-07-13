.PHONY: all build run clean swag test

all: build

build:
	@go build -o moneyger cmd/moneyger/main.go

run: swag build
	@./moneyger

clean:
	@rm -f moneyger

swag:
	@swag init -g cmd/moneyger/main.go --output docs

test:
	@go test -v ./...
