package domain

type WorkerStatus string

const (
	Success WorkerStatus = "success"
	Failure WorkerStatus = "failure"
)

type WorkerResult struct {
	Status WorkerStatus
	Url    string
}
