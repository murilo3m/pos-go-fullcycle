package load_testing

import (
	"cli/internal/flags"
	"fmt"
	"time"
)

type LoadTestingCmd struct{}

type Result struct {
	StatusCode int
	Duration   time.Duration
	Error      error
}

type LoadTestingFlags struct {
	URL         string
	Requests    int
	Concurrency int
}

func NewLoadTestingCmd() *LoadTestingCmd {
	return &LoadTestingCmd{}
}

func (c *LoadTestingCmd) Execute(parsedFlags flags.ParsedFlags) error {
	fmt.Println("Starting load testing")
	fmt.Printf("The URL to be used is %s, the quantity of requests is %d, and the concurrency is %d\n", parsedFlags.URL, parsedFlags.Requests, parsedFlags.Concurrency)

	flags := LoadTestingFlags{
		URL:         parsedFlags.URL,
		Requests:    parsedFlags.Requests,
		Concurrency: parsedFlags.Concurrency,
	}

	startTime := time.Now()

	results := RunLoadTest(flags)

	totalExecutionTime := time.Since(startTime)

	GenerateReport(flags, results, totalExecutionTime)

	return nil
}
