package main

import "fmt"

func main() {
	// ketika konversi tipe data jika ukurannya tidak cukup akan ada perubahan data (hati-hati)
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)

	// terjadi integer overflow, saat mencapai batas akan balik ke titik bawah
	// hati-hati saat lakukan konversi
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// byte alias dari uint8
	// konversi string jika ingin ambil karakter di string
	var name = "Dicki"
	var e = name[1]
	var eToString = string(e)
	fmt.Println(eToString)
}
