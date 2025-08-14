package flags

import (
	"errors"
	"flag"
	"fmt"
)

type ParsedFlags struct {
	URL         string
	Requests    int
	Concurrency int
}

func showHelp() error {
	helpText := `uso:
  	DEV: go run cmd/cli/main.go -u <URL> -r <N> -c <N>
	PROD "Executável" -u <URL> -r <N> -c <N>

	opções:
  		-u, --url          URL do serviço a ser testado (obrigatório)
  		-r, --requests     número total de requests (obrigatório)
 		-c, --concurrency  número de chamadas simultâneas (obrigatório)
  		-h, --help         mostra esta ajuda`
	return errors.New(helpText)
}

func ParseFlags() (ParsedFlags, error) {
	flag.Usage = func() {
		fmt.Println("erro: argumentos inválidos")
		fmt.Println(showHelp().Error())
	}

	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 0, "número total de requests")
	concurrency := flag.Int("concurrency", 0, "número de chamadas simultâneas")
	helpFlag := flag.Bool("help", false, "mostra ajuda")

	flag.StringVar(url, "u", "", "alias para --url")
	flag.IntVar(requests, "r", 0, "alias para --requests")
	flag.IntVar(concurrency, "c", 0, "alias para --concurrency")

	flag.Parse()

	if *helpFlag {
		return ParsedFlags{}, showHelp()
	}

	if *url == "" || *requests == 0 || *concurrency == 0 {
		return ParsedFlags{}, showHelp()
	}

	return ParsedFlags{
		URL:         *url,
		Requests:    *requests,
		Concurrency: *concurrency,
	}, nil
}
