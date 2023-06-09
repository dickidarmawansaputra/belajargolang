package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// package container/rint = implementasi struktur data circular list
	// circular list = struktur data ring, dimana diakhir element akan kembali ke element awal (HEAD)
	data := ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value: " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
