package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	serverURL      = "http://localhost:8080/cotacao"
	clientTimeout  = 300 * time.Millisecond
	outputFileName = "cotacao.txt"
)

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
	defer cancel()

	cotacao, err := getCotacaoFromServer(ctx)
	if err != nil {
		log.Printf("Erro ao obter cotação do servidor: %v", err)
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout excedido ao receber resultado do servidor.")
		}
		return
	}

	err = saveCotacaoToFile(cotacao.Bid)
	if err != nil {
		log.Printf("Erro ao salvar cotação no arquivo: %v", err)
		return
	}

	fmt.Printf("Cotação do Dólar salva com sucesso em '%s': Dólar: %s\n", outputFileName, cotacao.Bid)
}

func getCotacaoFromServer(ctx context.Context) (*CotacaoResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", serverURL, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição HTTP: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição HTTP para o servidor: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("resposta inesperada do servidor: %s", resp.Status)
	}

	var cotacao CotacaoResponse
	err = json.NewDecoder(resp.Body).Decode(&cotacao)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON do servidor: %w", err)
	}

	return &cotacao, nil
}

func saveCotacaoToFile(bid string) error {
	content := fmt.Sprintf("Dólar: %s", bid)
	err := os.WriteFile(outputFileName, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo '%s': %w", outputFileName, err)
	}
	return nil
}
