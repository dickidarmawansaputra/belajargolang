package main

func main() {
	// template function
	/*
		● Selain mengakses field, dalam template, function juga bisa diakses
		● Cara mengakses function sama seperti mengakses field, namun jika function tersebut memiliki
		parameter, kita bisa gunakan tambahkan parameter ketika memanggil function di template nya
		● {{.FunctionName}}, memanggil field FunctionName atau function FunctionName()
		● {{.FunctionName “eko”, “kurniawan”}}, memanggil function FunctionName(“eko”, “kurniawan”)
	*/

	// global function
	/*
		● Go-Lang template memiliki beberapa global function
		● Global function adalah function yang bisa digunakan secara langsung, tanpa menggunakan
		template data
		● Berikut adalah beberapa global function di Go-Lang template
		● https://github.com/golang/go/blob/master/src/text/template/funcs.go
	*/

	// Menambah Global Function
	/*
		● Kita juga bisa menambah global function
		● Untuk menambah global function, kita bisa menggunakan method Funcs pada template
		● Perlu diingat, bahwa menambahkan global function harus dilakukan sebelum melakukan parsing
		template
	*/

	// Function Pipelines
	/*
		● Go-Lang template mendukung function pipelines, artinya hasil dari function bisa dikirim ke
		function berikutnya
		● Untuk menggunakan function pipelines, kita bisa menggunakan tanda | , misal :
		● {{ sayHello .Name | upper }}, artinya akan memanggil global function sayHello(Name) hasil dari
		sayHello(Name) akan dikirim ke function upper(hasil)
		● Kita bisa menambahkan function pipelines lebih dari satu
	*/
}
