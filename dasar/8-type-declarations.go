package main

import "fmt"

func main() {
	// kemampunan membuat ulang tipe data baru dari tipe data yang sudah ada
	// atau disebut buat alias
	type NoKTP string

	// tipe data harus sesuai dengan type declarations yang telah dibuat
	var noKtp NoKTP = "123"
	fmt.Println(noKtp)
}
