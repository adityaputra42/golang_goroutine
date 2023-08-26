package golanggoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomacprox(t *testing.T) {

	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	totalCpu := runtime.NumCPU()
	fmt.Println("total cpu : ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total Thread", totalThread)

	totalGororutne := runtime.NumGoroutine()

	fmt.Println("total Goroutine : ", totalGororutne)
	
	group.Wait()
}
