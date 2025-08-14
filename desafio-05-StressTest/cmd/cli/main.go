package main

import (
	"cli/internal/dispatcher"
	"cli/internal/flags"
	"fmt"
	"os"
)

func main() {
	parsedFlags, err := flags.ParseFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = flags.ValidateFlags(parsedFlags)
	if err != nil {
		fmt.Printf("Flags inválidas: %v\n", err)
		os.Exit(1)
	}

	dispatcher.Dispatch(parsedFlags)
}
