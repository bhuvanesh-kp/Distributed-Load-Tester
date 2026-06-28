package domain

import "time"

type WorkerStatus string

const (
	Success WorkerStatus = "success"
	Failure WorkerStatus = "failure"
)

type WorkerResult struct {
	Status   WorkerStatus
	Url      string
	Duration time.Duration
}
