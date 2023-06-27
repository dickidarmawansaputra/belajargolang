package main

func main() {
	// middleware
	/*
		● Dalam pembuatan web, ada konsep yang bernama middleware atau filter atau interceptor
		● Middleware adalah sebuah fitur dimana kita bisa menambahkan kode sebelum dan setelah sebuah
		handler di eksekusi
	*/

	// Middleware di Go-Lang web
	/*
		● Sayangnya, di Go-Lang web tidak ada middleware
		● Namun karena struktur handler yang baik menggunakan interface, kita bisa membuat middleware
		sendiri menggunakan handler
	*/

	// Error Handler
	/*
		● Kadang middleware juga biasa digunakan untuk melakukan error handler
		● Hal ini sehingga jika terjadi panic di Handler, kita bisa melakukan recover di middleware, dan
		mengubah panic tersebut menjadi error response
		● Dengan ini, kita bisa menjaga aplikasi kita tidak berhenti berjalan
	*/
}
