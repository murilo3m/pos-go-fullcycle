package handlers

import (
	"cepgraus/internal/clients"
	"cepgraus/internal/services"
	"cepgraus/internal/utils"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type TemperatureHandler struct {
	CityClient        *clients.CityClient
	TemperatureClient *clients.TemperatureClient
}

func NewTemperatureHandler(cityClient *clients.CityClient, temperatureClient *clients.TemperatureClient) *TemperatureHandler {
	return &TemperatureHandler{
		CityClient:        cityClient,
		TemperatureClient: temperatureClient,
	}
}

func (th *TemperatureHandler) TemperatureByCep(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if !services.ValidateCep(cep) {
		http.Error(w, "Invalid zipcode", http.StatusBadRequest)
		return
	}

	cityName, err := th.CityClient.GetCityByCEP(cep)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Can not find zipcode", http.StatusNotFound)
		return
	}

	cityNameFormatted := utils.Formatter(cityName.Localidade)

	temperature, err := th.TemperatureClient.GetTemperature(cityNameFormatted)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Can not find temperature", http.StatusBadRequest)
		return
	}

	temperatureConverter, err := utils.NewTemperatureConverter(temperature, "C")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"temp_C": math.Round(temperatureConverter.ToCelsius()*10) / 10,
		"temp_F": math.Round(temperatureConverter.ToFahrenheit()*10) / 10,
		"temp_K": math.Round(temperatureConverter.ToKelvin()*10) / 10,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
