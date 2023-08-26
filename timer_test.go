package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C

	fmt.Println(time)

}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("time now", time.Now())

	time := <-channel

	fmt.Println("time after", time)

}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("dieksekusi setelah 1 detik")
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())
	group.Wait()
}
