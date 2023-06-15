package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRaceAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				// x = x + 1 semula seperti ini, diganti dengan atomic dan aman dari race condition
				// result counter = 100000
				// tapi jika datanya seperti struct cocok dengan mutex
				// atomic cocok untuk data primitive seperti angka
				// Tipe data primitif adalah tipe data bawaan dari bahasa pemrograman, seperti di Java ada delapan tipe data primitif, dan semuanya sudah kita bahas yaitu : byte. short. int.
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()

	time.Sleep(5 * time.Second)
	fmt.Println("counter = ", x)
}
