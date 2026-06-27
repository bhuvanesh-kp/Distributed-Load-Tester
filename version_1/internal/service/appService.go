package service

import "version_1/internal/domain"

type AppServiceInterface interface {
	LoadWithSingleWorker(url string) any
}

func NewService() AppServiceInterface {
	return &NewServiceStruct{}
}

type NewServiceStruct struct{}

func (srv *NewServiceStruct) LoadWithSingleWorker(url string) any {
	res := domain.WorkerResult{
		Status: domain.Success,
		Url:    url,
	}

	return res
}
