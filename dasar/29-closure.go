package main

import "fmt"

func main() {
	// closure = kemampuan sebuah func interaksi dengan data-data disekitarnya
	// gunakan fitur ini dengan bijak (hati-hati)
	counter := 0
	increment := func() {
		// scope diatas ini bisa diakses
		// hati-hati merubah data
		// solusi buat variable baru atau deklarasi ulang
		// jadi tidak merubah data diatasnya
		counter := 0
		fmt.Println("Increment")
		counter++
		fmt.Println("Hasil increment", counter)
	}

	increment()
	fmt.Println(counter)
}
