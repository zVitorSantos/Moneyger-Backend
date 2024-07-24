FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
COPY deployments/.env .env
COPY .env.example .env.example
RUN ls -la /app

RUN go build -o moneyger cmd/moneyger/main.go

EXPOSE 8080

CMD ["./moneyger"]
