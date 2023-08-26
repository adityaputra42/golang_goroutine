package golanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	chanel := make(chan string)

	defer close(chanel)

	go func() {
		time.Sleep(2 * time.Second)
		chanel <- "Gita Prigi Andani"
		fmt.Println("Selesai mengirim data")
	}()
	data := <-chanel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeresponse(chanel chan string) {
	time.Sleep(2 * time.Second)
	chanel <- "Gita Aditya Putra"
}

func TestChannelAsParameter(t *testing.T) {
	chanel := make(chan string)

	defer close(chanel)

	go GiveMeresponse(chanel)
	data := <-chanel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func in(chanel chan<- string) {
	time.Sleep(2 * time.Second)
	chanel <- "Gita Aditya Putra"
}

func out(chanel <-chan string) {
	data := <-chanel
	fmt.Println(data)
}

func TestChannelInOut(t *testing.T) {
	chanel := make(chan string)
	go in(chanel)
	go out(chanel)

	time.Sleep(3 * time.Second)
	close(chanel)
}

func TestBufferedChannel(t *testing.T) {
	chanel := make(chan string, 3)
	defer close(chanel)
	chanel <- "Gita"
	chanel <- "Prigi"
	chanel <- "Aditya"

	fmt.Println(<-chanel)

	fmt.Println(<-chanel)

	fmt.Println(<-chanel)
	fmt.Println("Seleasi")
}

func TestRangeChannel(t *testing.T) {

	chanel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			chanel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(chanel)
	}()

	for data := range chanel {
		fmt.Println("Menerima data => " + data)
	}
}

func TestSelectChanel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeresponse(channel1)
	go GiveMeresponse(channel2)
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChanel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeresponse(channel1)
	go GiveMeresponse(channel2)
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
