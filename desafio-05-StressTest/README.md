# Desafio - Pós Go Lang - Full Cycle

## Objetivo

Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

## Requisitos

- Docker.

## Configuração e Execução

1. **Build da imagem Docker:**

   No diretório raiz do projeto, execute:

   ```sh
   docker build -t "nome-da-imagem" .

   ```

2. **Execução do container:**

   No terminal, rode o seguinte comando:

   ```sh
   docker run "nome-da-imagem" --url=http://google.com --requests=1000 --concurrency=10 .

   ```

3. **Regras e Validações**

   - O sistema aceita as flags em formato curto:
     -u para --url
     -r para --requests
     -c para --concurrency

   - A flag url deve ser um endereço válido.
   - O número máximo de requisições permitidas (--requests) é 1000.
   - O número máximo de concorrência (--concurrency) é 10.
