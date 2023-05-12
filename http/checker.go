package http

import (
	"net/http"
	"time"
)

func CheckURL(url string) URLStatus {
	start := time.Now()
	response, err := http.Get(url)

	if err != nil {
		return URLStatus{
			StatusCode: 0,
			Elapsed:    0,
			IsHealthy:  false,
			Error:      err,
		}
	}

	defer response.Body.Close()
	elapsed := time.Since(start)
	isHealthy := IsStatusCodeHealth(response.StatusCode)

	return URLStatus{
		StatusCode: response.StatusCode,
		Elapsed:    elapsed,
		IsHealthy:  isHealthy,
		Error:      nil,
	}
}

func IsStatusCodeHealth(statusCode int) bool {
	return !(statusCode == http.StatusInternalServerError || statusCode == http.StatusBadGateway || statusCode == http.StatusServiceUnavailable)
}
