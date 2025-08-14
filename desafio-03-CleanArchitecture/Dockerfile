FROM golang:1.22.6 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/src/cmd/ordersystem

RUN go mod tidy

RUN go install github.com/ktr0731/evans@latest

CMD ["go", "run", "main.go", "wire_gen.go"]
