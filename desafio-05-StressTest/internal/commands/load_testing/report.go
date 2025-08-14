// report.go
package load_testing

import (
	"fmt"
	"strings"
	"time"
)

func GenerateReport(flags LoadTestingFlags, results []Result, totalExecutionTime time.Duration) {
	totalDuration := time.Duration(0)
	statusCounts := make(map[int]int)
	errorCount := 0

	for _, result := range results {
		if result.Error != nil {
			errorCount++
			continue
		}
		statusCounts[result.StatusCode]++
		totalDuration += result.Duration
	}

	successCount := flags.Requests - errorCount
	var avgDuration time.Duration
	if successCount > 0 {
		avgDuration = totalDuration / time.Duration(successCount)
	}

	completedRequests := len(results)
	totalRequests := flags.Requests
	progressPercentage := float64(completedRequests) / float64(totalRequests) * 100

	totalMinutes := int(totalExecutionTime / time.Minute)
	totalSeconds := int((totalExecutionTime % time.Minute) / time.Second)
	totalMilliseconds := int((totalExecutionTime % time.Second) / time.Millisecond)

	minutes := int(totalDuration / time.Minute)
	seconds := int((totalDuration % time.Minute) / time.Second)
	milliseconds := int((totalDuration % time.Second) / time.Millisecond)

	fmt.Println("\nRelatório de teste de carga:")
	fmt.Printf("Total de requisições: %d\n", flags.Requests)
	fmt.Printf("Requisições completadas: %d (%.1f%%)\n", completedRequests, progressPercentage)
	fmt.Printf("Requisições com erro: %d\n", errorCount)

	fmt.Printf("Tempo total de execução: %dm %ds %dms\n", totalMinutes, totalSeconds, totalMilliseconds)
	fmt.Printf("Tempo acumulado das requisições: %dm %ds %dms\n", minutes, seconds, milliseconds)

	if successCount > 0 {
		fmt.Printf("Tempo médio por requisição: %dms\n", avgDuration.Milliseconds())
	}

	fmt.Println("\nDistribuição por status:")
	for status, count := range statusCounts {
		fmt.Printf("  Status %d: %d requisições (%.1f%%)\n",
			status,
			count,
			float64(count)/float64(flags.Requests)*100,
		)
	}
}

func ReportProgress(completedRequests, totalRequests int) {
	progressPercentage := float64(completedRequests) / float64(totalRequests) * 100

	fmt.Printf("\r")
	fmt.Printf("Progresso: [")
	width := 30
	completed := width * int(progressPercentage) / 100

	bar := strings.Repeat("█", completed) + strings.Repeat(" ", width-completed)
	fmt.Printf("%s] %d/%d (%.1f%%)", bar, completedRequests, totalRequests, progressPercentage)
}