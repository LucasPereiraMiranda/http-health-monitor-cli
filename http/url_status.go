package http

import "time"

type URLStatus struct {
	StatusCode int
	Elapsed    time.Duration
	IsHealthy  bool
	Error      error
}
