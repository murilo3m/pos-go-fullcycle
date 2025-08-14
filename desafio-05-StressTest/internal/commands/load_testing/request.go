package load_testing

import (
	"net/http"
	"time"
)

func MakeRequest(url string) Result {
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	if err != nil {
		return Result{StatusCode: 0, Duration: duration, Error: err}
	}
	defer resp.Body.Close()

	return Result{StatusCode: resp.StatusCode, Duration: duration, Error: nil}
}
