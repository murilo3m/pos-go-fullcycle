# Desafio - Pós Go Lang - Full Cycle

## Objetivo

Implementar uma API em go que recebe um CEP retorne a temperatura em Celsius, Kelvin e Fahrenheit.

## Funcionalidades

- Endpoint capaz de receber um CEP
- Retorna erro quando não recebeu CEP válido
- Quando não encontra a cidade
- Quando não encontra a temperatura da cidade

## Requisitos

- **APIKEY do weatherapi**: https://www.weatherapi.com/

## Configuração

1. Na raiz do projeto no arquivo crie .env usando o .env.example

   - Preencher com sua key do weatherapi

2. Acesse o diretório app

   - cd cmd/app
   - go run main.go

3. Use `http://localhost:8080/temperatureByCEP?cep=SEU CEP`:

   - Substitua com o seu CEP
   - Só é valido CEP com numeros

4. OBS

   - Existe o arquivo .http que faz a chamada só é necessario a extensão REST Client
   - Use o comando go test ./... para executar os testes

5. O link para acessar a API

   - `https://pos-go-fullcycle-weatherbycep-541049047259.us-central1.run.app/temperatureByCEP?cep=14401220`
   - OBS: Vai ser desativada após a avaliação
