package main

import "fmt"

func main() {
	// break untuk menghentikan seluruh perulangan
	// continue untuk menghentikan perulangan yang berjalan, dan lanjut perulangan selanjutnya
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke", i)
	}

	// angka ganjil
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke", i)
	}
}
