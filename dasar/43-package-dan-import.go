package main

import (
	"dasar/dasar/helper"
	"fmt"
)

func main() {
	/**
	Error Package dan Importt
	Di Go-Lang versi terbaru, kita sudah diwajibkan menggunakan Go-Modules, yang akan dijelaskan pada PART 2 kelas ini

	Jadi pada video selanjutnya, pasti akan mendapat error

	Namun jika ingin mencoba fitur lama Go-Lang, sebelum lanjutkan video selanjutnya, silahkan matikan default fitur Go-Modules dengan menggunakan perintah

	go env -w  GO111MODULE=off
	*/

	// package = tempat yg bisa digunakan untuk mengorganisir kode program ygn kita buat di golang
	// dengan package bisa merapikan kode program
	// package sendiri di go hanya direktori folder di sistem operasi
	result := helper.SayHello("Dicki")
	fmt.Println(result)
}
