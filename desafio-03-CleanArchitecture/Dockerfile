# Estágio de build
FROM golang:1.23.4-alpine AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos necessários para o build
COPY . .

# Instala a CLI do migrate
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Instala dependências e faz o build da aplicação
RUN cd cmd/ordersystem && go mod tidy && go build -o ordersystem .

# Estágio de execução
FROM golang:1.23.4-alpine

# Define o diretório de trabalho
WORKDIR /app

# Instala dependências necessárias (bash e curl para baixar o wait-for)
RUN apk add --no-cache bash curl

# Baixa o script wait-for diretamente do repositório oficial
RUN curl -o /wait-for https://raw.githubusercontent.com/eficode/wait-for/master/wait-for && \
    chmod +x /wait-for

# Copia o binário da aplicação, a CLI do migrate e as migrations
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/cmd/ordersystem/ordersystem .
COPY --from=builder /app/cmd/ordersystem/.env .
COPY --from=builder /app/internal/infra/database/migrations /migrations

# Expõe a porta da aplicação
EXPOSE 8080

# Comando para aguardar o MySQL e RabbitMQ estarem prontos, rodar as migrations e iniciar a aplicação
CMD sh -c "/wait-for mysql:3306 -- migrate -path /migrations -database 'mysql://root:root@tcp(mysql:3306)/orders' -verbose up && /wait-for rabbitmq:5672 -- ./ordersystem"