package golanggoroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello Adit")
}

func TestGoroutineCheck(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestDisplay(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}
