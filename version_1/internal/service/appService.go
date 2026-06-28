package service

import (
	"net/http"
	"time"
	"version_1/internal/domain"
)

type AppServiceInterface interface {
	LoadWithSingleWorker(url string) any
}

func NewService() AppServiceInterface {
	return &NewServiceStruct{}
}

type NewServiceStruct struct{}

func (srv *NewServiceStruct) LoadWithSingleWorker(url string) any {
	startTime := time.Now()

	_, err := http.Get(url)

	duration := time.Since(startTime)

	if err != nil {
		return domain.WorkerResult{
			Status: domain.Failure,
			Url:    url,
		}
	}

	res := domain.WorkerResult{
		Status:   domain.Success,
		Url:      url,
		Duration: time.Duration(duration.Microseconds()),
	}

	return res
}
