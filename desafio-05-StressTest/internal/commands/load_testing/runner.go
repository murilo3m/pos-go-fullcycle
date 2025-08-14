package load_testing

import (
	"sync"
	"sync/atomic"
)

func RunLoadTest(flags LoadTestingFlags) []Result {
	results := make(chan Result, flags.Requests)
	semaphore := make(chan struct{}, flags.Concurrency)
	var wg sync.WaitGroup

	var completedRequests int32

	for i := 0; i < flags.Requests; i++ {
		wg.Add(1)
		semaphore <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-semaphore }()

			result := MakeRequest(flags.URL)
			results <- result

			newCompleted := atomic.AddInt32(&completedRequests, 1)

			ReportProgress(int(newCompleted), flags.Requests)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var collectedResults []Result
	for result := range results {
		collectedResults = append(collectedResults, result)
	}

	return collectedResults
}
