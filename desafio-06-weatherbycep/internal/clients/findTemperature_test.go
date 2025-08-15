package clients

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var apiKey string = "kjsadh23y7896yff"

func TestGetTemperature_Success(t *testing.T) {
	response := `{
		"current": {
			"temp_c": 3.2
		}
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer ts.Close()

	client := NewTemperatureClient(apiKey)
	client.BaseURL = ts.URL

	temperature, err := client.GetTemperature("London")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedTemperature := 3.2
	if temperature != expectedTemperature {
		t.Fatalf("expected %.1f, got %.1f", expectedTemperature, temperature)
	}
}

func TestGetTemperature_CityNotFound(t *testing.T) {
	response := `{
		"error": {
			"code": 1006,
			"message": "No matching location found."
		}
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(response))
	}))
	defer ts.Close()

	client := NewTemperatureClient(apiKey)
	client.BaseURL = ts.URL

	_, err := client.GetTemperature("InvalidCity")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	expectedError := "API error: No matching location found. (code: 1006)"
	if err.Error() != expectedError {
		t.Fatalf("expected error %q, got %q", expectedError, err.Error())
	}
}
