package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"cepgraus/internal/clients"
	"cepgraus/internal/routes"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading file .env:", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not found in the environment")
	}

	cityClient := clients.NewCityClient()
	temperatureClient := clients.NewTemperatureClient(apiKey)

	router := routes.SetupRoutes(cityClient, temperatureClient)

	fmt.Println("Server start at http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Erro ao iniciar o servidor", err)
	}
}
