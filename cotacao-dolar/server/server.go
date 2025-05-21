package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	awesomeAPIURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	dbTimeout     = 10 * time.Millisecond
	apiTimeout    = 200 * time.Millisecond
)

type AwesomeAPIData struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

func main() {
	db, err := sql.Open("sqlite3", "./cotacao.db")
	if err != nil {
		log.Fatal("Erro ao abrir o banco de dados:", err)
	}
	defer db.Close()

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS cotacoes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        bid TEXT,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Erro ao criar a tabela:", err)
	}

	http.HandleFunc("/cotacao", cotacaoHandler(db))

	fmt.Println("Servidor iniciado na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func cotacaoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxAPI, cancelAPI := context.WithTimeout(r.Context(), apiTimeout)
		defer cancelAPI()

		cotacao, err := getCotacaoFromAwesomeAPI(ctxAPI)
		if err != nil {
			log.Printf("Erro ao obter cotação da API: %v", err)
			if ctxAPI.Err() == context.DeadlineExceeded {
				http.Error(w, "Timeout ao chamar a API de cotação", http.StatusRequestTimeout)
				return
			}
			http.Error(w, "Erro ao obter cotação", http.StatusInternalServerError)
			return
		}

		ctxDB, cancelDB := context.WithTimeout(r.Context(), dbTimeout)
		defer cancelDB()

		err = saveCotacaoToDB(ctxDB, db, cotacao.USDBRL.Bid)
		if err != nil {
			log.Printf("Erro ao salvar cotação no banco de dados: %v", err)
			if ctxDB.Err() == context.DeadlineExceeded {
				http.Error(w, "Timeout ao persistir dados no banco", http.StatusRequestTimeout)
				return
			}
			http.Error(w, "Erro ao salvar cotação", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(CotacaoResponse{Bid: cotacao.USDBRL.Bid})
	}
}

func getCotacaoFromAwesomeAPI(ctx context.Context) (*AwesomeAPIData, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", awesomeAPIURL, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição HTTP: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição HTTP para AwesomeAPI: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("resposta inesperada da AwesomeAPI: %s", resp.Status)
	}

	var data AwesomeAPIData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON da AwesomeAPI: %w", err)
	}

	return &data, nil
}

func saveCotacaoToDB(ctx context.Context, db *sql.DB, bid string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		stmt, err := db.PrepareContext(ctx, "INSERT INTO cotacoes(bid) VALUES(?)")
		if err != nil {
			return fmt.Errorf("erro ao preparar statement SQL: %w", err)
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(ctx, bid)
		if err != nil {
			return fmt.Errorf("erro ao executar insert no banco de dados: %w", err)
		}
		return nil
	}
}
