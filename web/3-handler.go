package main

func main() {
	// handler
	/*
		● Server hanya bertugas sebagai Web Server, sedangkan untuk menerima HTTP Request yang
		masuk ke Server, kita butuh yang namanya Handler
		● Handler di Go-Lang di representasikan dalam interface, dimana dalam kontraknya terdapat sebuah
		function bernama ServeHTTP() yang digunakan sebagai function yang akan di eksekusi ketika
		menerima HTTP Request
	*/

	// HandlerFunc
	/*
		● Salah satu implementasi dari interface Handler adalah HandlerFunc
		● Kita bisa menggunakan HandlerFunc untuk membuat function handler HTTP
	*/
}
