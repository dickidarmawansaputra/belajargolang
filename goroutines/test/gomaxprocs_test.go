package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}

	// ngecek jumlah goroutine
	// output 102
	// yg 2 satu untuk jalanin program 1 lagi untuk garbage collector
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("total CPU:", totalCpu)

	// kenap minus, karna kalo di atas 0 akan mengubah jumlah threadnya
	// jadi harus dibawah 0
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine:", totalGoroutine)

	group.Wait()
}
