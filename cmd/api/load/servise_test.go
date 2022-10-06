package load

import (
	"fmt"
	"sync"
	"testing"
)

var (
	wg          = &sync.WaitGroup{}
	RoutinCalls = 3
	racerNumberPerRoutineCall = 10
	loc = &sync.Mutex{}
)

func TestLoader(t *testing.T) {
	serv := NewloadService(NewloadRepo())
	wg.Add(RoutinCalls)
	mod4 := 0
	for i := 1; i <= RoutinCalls; i++ {
		go func(i int) {
			for j := 1; j <= racerNumberPerRoutineCall; j++ {
				fmt.Println("Calling Goroutine ---", i)
				if j % 4 == 0 {
					loc.Lock()
					mod4 ++
					loc.Unlock()
					//wrong url to get 401
					_, _ = serv.Loader("https://jsonplaceholder.typicode.com/todos/1s")
				}else{
					_, _ = serv.Loader("")
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	res, _ := serv.GetInfo()
	fmt.Println("resp is", res, mod4)
	if len(res) == 0 {
		fmt.Println("the length of the response must be greator than 0")
	}
	if res[404] != mod4{
		fmt.Println("the length of the failed request is of 404", mod4)
	}
}
func TestGetInfo(t *testing.T) {
	serv := NewloadService(NewloadRepo())
	_, _ = serv.Loader("")
	res, err := serv.GetInfo()
	if err != nil {
		fmt.Println("Err should be nil")
	}
	if len(res) == 0 {
		fmt.Println("the length of the response must be greator than 0")
	}
}
