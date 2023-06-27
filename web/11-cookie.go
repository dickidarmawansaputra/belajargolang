package main

func main() {
	// stateless
	/*
		● HTTP merupakan stateless antara client dan server, artinya server tidak akan menyimpan data
		apapun untuk mengingat setiap request dari client
		● Hal ini bertujuan agar mudah melakukan scalability di sisi server
		● Lantas bagaimana caranya agar server bisa mengingat sebuah client? Misal ketika kita sudah login
		di website, server otomatis harus tahu jika client tersebut sudah login, sehingga request
		selanjutnya, tidak perlu diminta untuk login lagi
		● Untuk melakukan hal ini, kita bisa memanfaatkan Cookie
	*/

	// cookie
	/*
		● Cookie adalah fitur di HTTP dimana server bisa memberi response cookie (key-value) dan client
		akan menyimpan cookie tersebut di web browser (otomatis menerima & menyimpan di browser)
		● Request selanjutnya, client akan selalu membawa cookie tersebut secara otomatis
		● Dan server secara otomatis akan selalu menerima data cookie yang dibawa oleh client setiap kalo
		client mengirimkan request
	*/

	// membuat cookie
	/*
		● Cookie merupakan data yang dibuat di server dan sengaja agar disimpan di web browser
		● Untuk membuat cookie di server, kita bisa menggunakan function http.SetCookie()
	*/
}
