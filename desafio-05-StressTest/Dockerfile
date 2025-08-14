FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/myapp ./cmd/cli

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

RUN adduser -D appuser

WORKDIR /app

COPY --from=builder /app/myapp .

RUN chown -R appuser:appuser /app

USER appuser

ENTRYPOINT ["/app/myapp"]

CMD []