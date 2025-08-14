# Desafio 3 Pós Go Lang - Full Cycle

## Objetivo

O objetivo principal foi criar um sistema utilizando à arquitetura limpa utilizando a linguagem go, o sistem permite criar e listar "orders" usando três tecnologias de comunicação(REST, GRPC e GRAPHQL)

## Funcionalidades

- É possivel criar uma order
- É possivel listar as orders

## Requisitos

- **Banco de dados**: MySQL
- **Mensageria**: RabbitMQ
- **Ferramentas Necessárias**:
  - Migration para criar o banco de dados "orders" e sua tabela. (Somente se usar o devcontainer)
  - Evans instalado para usar como cliente do gRPC. (Somente se usar o devcontainer)
  - Extensão Rest Client instalada no VS Code.

## Configuração

1. No diretório `src/cmd/ordersystem`:

   - Crie um arquivo chamado `.env`.
   - Copie todo o conteúdo do arquivo `env_example.txt`.
   - Preencha os campos `DB_HOST` e `RABBITMQ_HOST`.
   - **Observação**: Se estiver executando com o DevContainer, não há necessidade de alterar esses campos.

2. No arquivo `src/Makefile`:
   - Preencha o campo `DB_HOST`.
   - **Observação**: Se estiver executando com o DevContainer, não há necessidade de alterar esse campo.

## Executando o Projeto

### Com o DevContainer

- É necessário ter o Docker instalado e a extensão Dev Containers no VS Code.
- Ao abrir o diretório, o VS Code oferecerá a opção de reabrir no DevContainer.
- Ou use o atalho para abrir todos os comandos: `Ctrl + Shift + P` ou `Cmd + Shift + P` no macOS.
- Procure o comando "Reabrir diretório no container".

### Executando o Projeto Manualmente

1. No diretório `src`:

   - Execute o comando `make migrate` para rodar as migrações do banco de dados.
   - Execute o comando `go mod tidy` para garantir que todas as dependências estejam instaladas.
   - Execute o comando `go run main.go wire_gen.go` para iniciar a aplicação.

2. No diretório `src/api`:

   - Utilize os dois arquivos para fazer chamadas REST (criar e listar pedidos).

3. No navegador, acesse `localhost:8080`:

   - Utilize o playground do GraphQL para realizar as operações:

   ```graphql
   mutation createOrder {
     createOrder(input: { id: "4", Price: 1287.00, Tax: 111.99 }) {
       id
       Price
       Tax
       FinalPrice
     }
   }

   query listOrders {
     orders {
       id
       Price
       Tax
       FinalPrice
     }
   }
   ```

4. No terminal
   - Execute o comando `evans -r repl` para iniciar o EVANS.
   - Execute o comando `package pb` para selecionar o pacote.
   - Execute o comando `service OrderService`.
   - Execute o comando `call CreateOrder` para iniciar a chamada de criação.
   - Execute o comando `call ListOrders` para iniciar a chamada de listagem.

OBS: Use o docker-compose da raiz do projeto para ter o mysql já configurado sem a necessidade de usar o comando make.
OBS: O .env já esta configurado, ao usar docker compose up, o projeto já vai estar funcionando na porta 8080(GRAPHQL) e 8000(REST)
OBS: Acesse o container do app para usar o EVANS

docker compose up
