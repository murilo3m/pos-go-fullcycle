FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM scratch

WORKDIR /app

COPY --from=builder /app/.env .

WORKDIR /app/cmd/app/

COPY --from=builder /app/cmd/app/main .

CMD ["./main"]
