package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	// hati - hati lupa menambahkan Done() karna akan berjalan terus tanpa berhenti
	defer group.Done()

	// akan me-running 1 proses asynchronous
	group.Add(1)

	fmt.Println("hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	// tidak perlu pake time.Sleep() untuk menunggu semua proses selesai
	fmt.Println("selesai")
}
