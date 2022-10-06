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
	// Tester() (string, httperrors.HttpErr)
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

// func (r *loadrepository) Tester() (string, httperrors.HttpErr) {
// 	serv := NewloadService(NewloadRepo())
// 	Wg.Add(routinCalls)
// 	mod4 := 0
// 	for i := 1; i <= routinCalls; i++ {
// 		go func(i int) {
// 			for j := 1; j <= RacerNumberPerRoutineCall; j++ {
// 				fmt.Println("Calling Goroutine ---", i)
// 				if j%4 == 0 {
// 					locker.Lock()
// 					mod4++
// 					locker.Unlock()
// 					//wrong url to get 401
// 					_ = r.loading("https://jsonplaceholder.typicode.com/todos/1s")
// 				} else {
// 					_ = r.loading("")
// 				}
// 			}
// 			Wg.Done()
// 		}(i)
// 	}
// 	Wg.Wait()
// 	res, _ := serv.GetInfo()
// 	fmt.Println("resp is", res, mod4)
// 	if len(res) == 0 {
// 		fmt.Println("the length of the response must be greator than 0")
// 	}
// 	if res[404] != mod4 {
// 		fmt.Println("the length of the failed request is of 404", mod4)
// 	}
// 	return "goroutines spanned", nil
// }