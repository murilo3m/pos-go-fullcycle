package flags

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func ValidateFlags(flags ParsedFlags) error {
	var validationErrors []string

	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(flags.URL)
	if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 400 {
		validationErrors = append(validationErrors, "URL inválida ou inacessível")
	}

	if flags.Requests > 1000 {
		validationErrors = append(validationErrors, "O número de requests não pode ser maior que 1000")
	}

	if flags.Concurrency > 10 {
		validationErrors = append(validationErrors, "O número de chamadas simultâneas (concurrency) não pode ser maior que 10")
	}

	if len(validationErrors) > 0 {
		return fmt.Errorf("erros de validação:\n- %s", strings.Join(validationErrors, "\n- "))
	}

	return nil
}
