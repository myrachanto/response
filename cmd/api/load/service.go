package load

import (
	httperrors "github.com/myrachanto/erroring"
)

var (
	LoadService LoadServiceInterface = &loadService{}
)

type LoadServiceInterface interface {
	Loader(url string) (*Load, httperrors.HttpErr)
	GetInfo() (map[int]int, httperrors.HttpErr)
}
type loadService struct {
	repo LoadRepoInterface
}

func NewloadService(repository LoadRepoInterface) LoadServiceInterface {
	return &loadService{
		repository,
	}
}
func (service *loadService) Loader(url string) (*Load, httperrors.HttpErr) {
	load, err := service.repo.Loader(url)
	return load, err
}

func (service *loadService) GetInfo() (map[int]int, httperrors.HttpErr) {
	loads, err := service.repo.GetInfo()
	return loads, err
}
