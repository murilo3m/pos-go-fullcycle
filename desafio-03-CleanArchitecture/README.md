# Order System Setup

Este guia fornece instruções para configurar e executar o Order System.

## Pré-requisitos

### Evans (Cliente gRPC)

Para instalar o Evans (cliente gRPC), siga as instruções em:
https://github.com/ktr0731/evans?tab=readme-ov-file#installation

## Iniciar o Projeto

1. Inicie os containers Docker na raiz do projeto:
```
docker compose up --build -d
```

## Testando o Projeto

### GraphQL
- Acesse: http://localhost:8080/

### gRPC
1. Inicie o Evans em modo REPL:
```
evans -r repl
```

2. Configure o ambiente:
```bash
package pb
service OrderService
```

3. Execute os comandos disponíveis:
```bash
call <criar/listar>
```

### api REST
1. Execute os comandos no arquivo `api/api.http`

### RabbitMQ
1. Acesse o painel do RabbitMQ:
   - URL: http://localhost:15672
   - Usuário: guest
   - Senha: guest

2. Configure a mensageria:
   - Crie uma nova fila chamada `orders`
   - Faça o bind da fila com a exchange `amq.direct`