package load

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	Loadrepository LoadRepoInterface = &loadrepository{}
	Repo                             = loadrepository{}
)

type LoadRepoInterface interface {
	GetURL(urls []string) []Load
	loading(urls []string) <-chan Load
}
type loadrepository struct {
	Store []Load `json:"store,omitempty"`
}

func NewloadRepo() LoadRepoInterface {
	return &loadrepository{}
}

func (r *loadrepository) GetURL(urls []string) []Load {
	results := r.loading(urls)
	for c := range results {
		r.Store = append(r.Store, c)
	}
	fmt.Println("-------------sdsff", r.Store)
	resp := sort(r.Store)
	return resp

}

func (r *loadrepository) loading(urls []string) <-chan Load {
	// errs := make(chan httperrors.HttpErr)
	results := make(chan Load)
	wg := &sync.WaitGroup{}
	goroutines := len(urls)
	wg.Add(goroutines)
	for _, url := range urls {
		go func(url string) {
			resp := Load{}
			if url == "" {
				// errs <- httperrors.NewBadRequestError("URL is empty")
				return
			}
			response, err := http.Get(url)
			if err != nil {
				// errs <- httperrors.NewBadRequestError("Something wrong with fetching the url")
				return
			}
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				// errs <- httperrors.NewBadRequestError("Something Went wrong with reading the response")
				return
			}
			size := len(string(body))
			resp.Size = size
			resp.Url = url
			results <- resp

			wg.Done()
		}(url)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}
func sort(data []Load) []Load {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i].Size > data[j].Size {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}
