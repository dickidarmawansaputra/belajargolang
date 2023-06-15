package test

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello world!")
}

func TestCreateGoroutine(t *testing.T) {
	// tidak akan tunggu RunHelloWorld() selesai dulu karna goroutine berjalan secara asynchronous
	go RunHelloWorld()
	fmt.Println("Ups")

	// memastika goroutine selesai dieksekusi (kasi jeda waktu 1 detik)
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Print("Display", number)
}

// membuktikan goroutine sangat ringan
func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
