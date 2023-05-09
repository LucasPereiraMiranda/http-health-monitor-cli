package http

import (
	"net/http"
	"time"
)

func CheckURL(url string) (int, time.Duration, bool) {
	start := time.Now()
	response, err := http.Get(url)

	if err != nil {
		return 0, 0, false
	}
	defer response.Body.Close()
	elapsed := time.Since(start)
	isHealthy := IsStatusCodeHealth(response.StatusCode)
	return response.StatusCode, elapsed, isHealthy
}

func IsStatusCodeHealth(statusCode int) bool {
	return !(statusCode == http.StatusInternalServerError || statusCode == http.StatusBadGateway || statusCode == http.StatusServiceUnavailable)
}
