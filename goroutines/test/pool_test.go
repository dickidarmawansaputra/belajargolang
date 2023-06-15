package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		// tambahkan ini jika butuh default value pool jika nil
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Dicki")
	pool.Put("Darmawan")
	pool.Put("Saputra")

	for i := 0; i < 10; i++ {
		go func() {
			// mengambil data dari pool (tidak berurutan ambil datanya, secara random)
			data := pool.Get()
			fmt.Println(data)
			// menambahkan datanya kembali ke dalam pool, karna ketika di get datanya maka akan kosong
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("selesai")
}
