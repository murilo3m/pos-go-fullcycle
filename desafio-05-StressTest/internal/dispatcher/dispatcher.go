package dispatcher

import (
	"cli/internal/commands/load_testing"
	"cli/internal/flags"
	"fmt"
	"os"
)

func Dispatch(parsedFlags flags.ParsedFlags) {
	cmd := load_testing.NewLoadTestingCmd()

	err := cmd.Execute(parsedFlags)
	if err != nil {
		fmt.Printf("Erro ao executar o benchmark: %v\n", err)
		os.Exit(1)
	}
}
