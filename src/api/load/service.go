package load

import (
	"strings"
)

var (
	LoadService LoadServiceInterface = &loadService{}
)

type LoadServiceInterface interface {
	GetURL(urls string) []Load
}
type loadService struct {
	repo LoadRepoInterface
}

func NewloadService(repository LoadRepoInterface) LoadServiceInterface {
	return &loadService{
		repository,
	}
}
func (service *loadService) GetURL(url string) []Load {
	if strings.Contains(url, ",") {
		urls := strings.Split(url, ",")
		// return Repo.GetURL(urls)
		return service.repo.GetURL(urls)
	}
	urls := strings.Split(url, " ")
	// return Repo.GetURL(urls)
	return service.repo.GetURL(urls)
}
