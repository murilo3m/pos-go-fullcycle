package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Endereco struct {
	Cep        string
	Logradouro string
	Bairro     string
	Localidade string
	Uf         string
	ApiOrigem  string
}

type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type ViaCEPResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func buscaBrasilAPI(ctx context.Context, cep string, ch chan<- Endereco) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var b BrasilAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&b); err != nil {
		return
	}
	e := Endereco{
		Cep:        b.Cep,
		Logradouro: b.Street,
		Bairro:     b.Neighborhood,
		Localidade: b.City,
		Uf:         b.State,
		ApiOrigem:  "BrasilAPI",
	}
	ch <- e
}

func buscaViaCEP(ctx context.Context, cep string, ch chan<- Endereco) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var v ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return
	}
	e := Endereco{
		Cep:        v.Cep,
		Logradouro: v.Logradouro,
		Bairro:     v.Bairro,
		Localidade: v.Localidade,
		Uf:         v.Uf,
		ApiOrigem:  "ViaCEP",
	}
	ch <- e
}

func main() {
	cep := "14401220"
	ch := make(chan Endereco, 2)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go buscaBrasilAPI(ctx, cep, ch)
	go buscaViaCEP(ctx, cep, ch)

	select {
	case res := <-ch:
		fmt.Printf("Resposta recebida da %s:\n", res.ApiOrigem)
		fmt.Printf("CEP: %s\nLogradouro: %s\nBairro: %s\nCidade: %s\nUF: %s\n",
			res.Cep, res.Logradouro, res.Bairro, res.Localidade, res.Uf)
	case <-ctx.Done():
		fmt.Println("Erro: Timeout de 1 segundo atingido.")
	}
}
