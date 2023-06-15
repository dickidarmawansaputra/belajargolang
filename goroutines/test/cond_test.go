package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	// jika melakukan wait pastikan ada yang mengirim signal misal cond.Signal()
	// jika tanpa mengirim signal, maka waitnya tidak akan pernah selesai karna tidak diberitahu goroutine untuk jalan / tidak perlu tunggu lagi
	cond.Wait()

	fmt.Println("done", value)

	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			// tiap habis 1 detik, eksekusi satu-satu
			cond.Signal()
		}
	}()

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Broadcast()
	// 	}
	// }()

	group.Wait()
}
