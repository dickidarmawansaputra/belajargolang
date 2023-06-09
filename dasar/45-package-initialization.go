package main

import _ "dasar/dasar/database" // tanda _ itu blank identifier

func main() {
	// saat membuat package, bisa membuat sebuah func yang akan diakses ketika package diakses
	// sangat cocok jika package berisi func2 untuk berkomunikasi dengan database
	// untuk membuat func yg diakses secara otomatis ketika package diakses cukup membuat func dengan nama init
	
	// di comment untuk test blank identifier
	// db := database.GetDatabase()
	// fmt.Println(db) 

	// blank identifier
	// kadang hanya ingin menjalankan init func di package tanpa harus eksekusi salah satu func di package
	// secara default, di go akan komplen ketika ada package yg diimport tapi tidak digunakan
	// untuk menangani ini, bisa menggunakan blank identifier dengan tanda _ sebelum nama package ketika import

}
