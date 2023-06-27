package main

func main() {
	// header
	/*
		● Selain Query Parameter, dalam HTTP, ada juga yang bernama Header
		● Header adalah informasi tambahan yang biasa dikirim dari client ke server atau sebaliknya
		● Jadi dalam Header, tidak hanya ada pada HTTP Request, pada HTTP Response pun kita bisa
		menambahkan informasi header
		● Saat kita menggunakan browser, biasanya secara otomatis header akan ditambahkan oleh
		browser, seperti informasi browser, jenis tipe content yang dikirim dan diterima oleh browser, dan
		lain-lain
	*/

	// request header
	/*
		● Untuk menangkap request header yang dikirim oleh client, kita bisa mengambilnya di
		Request.Header
		● Header mirip seperti Query Parameter, isinya adalah map[string][]string
		● Berbeda dengan Query Parameter yang case sensitive, secara spesifikasi, Header key tidaklah case
		sensitive
	*/

	// response header
	/*
		● Sedangkan jika kita ingin menambahkan header pada response, kita bisa menggunakan function
		ResponseWriter.Header()
	*/
}
