package main

func main() {
	// form post
	/*
		● Saat kita belajar HTML, kita tahu bahwa saat kita membuat form, kita bisa submit datanya dengan
		method GET atau POST
		● Jika menggunakan method GET, maka hasilnya semua data di form akan menjadi query parameter
		● Sedangkan jika menggunakan POST, maka semua data di form akan dikirim via body HTTP request
		● Di Go-Lang, untuk mengambil data Form Post sangatlah mudah
	*/

	// Request.PostForm
	/*
		● Semua data form post yang dikirim dari client, secara otomatis akan disimpan dalam attribute
		Request.PostForm
		● Namun sebelum kita bisa mengambil data di attribute PostForm, kita wajib memanggil method
		Request.ParseForm() terlebih dahulu, method ini digunakan untuk melakukan parsing data body
		apakah bisa di parsing menjadi form data atau tidak, jika tidak bisa di parsing, maka akan
		menyebabkan error
	*/
}
