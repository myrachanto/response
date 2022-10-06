package load

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	httperrors "github.com/myrachanto/erroring"
)

var (
	Loadrepository LoadRepoInterface = &loadrepository{}
	locker                           = &sync.Mutex{}
	Store          map[int]int       = make(map[int]int)
)

type LoadRepoInterface interface {
	Loader(url string) (*Load, httperrors.HttpErr)
	GetInfo() (map[int]int, httperrors.HttpErr)
	loading(url string) *Load
}
type Response struct {
	UserId    int    `json:"userId,omitempty"`
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

type loadrepository struct {
}

func NewloadRepo() LoadRepoInterface {
	return &loadrepository{}
}

func (r *loadrepository) Loader(url string) (*Load, httperrors.HttpErr) {
	l := r.loading(url)
	return l, nil

}
func (r *loadrepository) GetInfo() (map[int]int, httperrors.HttpErr) {
	return Store, nil
}

func (r *loadrepository) loading(url string) *Load {
	l := &Load{}
	resp := Response{}
	if url == "" {
		url = "https://jsonplaceholder.typicode.com/todos/1"
	}
	response, err := http.Get(url)
	if err != nil {
		locker.Lock()
		Store[response.StatusCode]++
		l.ResponseCode = response.StatusCode
		locker.Unlock()
		return l
	}
	defer response.Body.Close()
	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		locker.Lock()
		l.ResponseCode = 501
		Store[l.ResponseCode]++
		l.ResponseCode = response.StatusCode
		locker.Unlock()
		return l
	}
	err2 := json.Unmarshal([]byte(body), &resp)
	if err2 != nil {
		locker.Lock()
		l.ResponseCode = 500
		Store[l.ResponseCode]++
		l.ResponseCode = response.StatusCode
		locker.Unlock()
		return l
	}
	locker.Lock()
	fmt.Println(response.StatusCode)
	Store[response.StatusCode]++
	l.ResponseCode = response.StatusCode
	locker.Unlock()
	return l
}
