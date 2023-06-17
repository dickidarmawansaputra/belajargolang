package main

func main() {
	// pengenalan context
	/*
		context = sebuah data yg membawa value, sinyal cancel, sinyal timeout dan sinyal deadline
		context biasanya dibuat per request (misal setiap ada request masuk ke server web melalui http request)
		context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar proses
	*/

	// kenapa context perlu dipelajari?
	/*
		context di golang biasa digunakan untuk mengirim data request atau sinyal ke proses lain
		dengan menggunakan context, ketika ingin membatalkan semua proses, cukup mengirim sinyal ke context, maka secara otomatis semua proses akan dibatalkan
		hampir semua bagian di golang memanfaatkan context, seperti database, http server, http client, dan lainnya
		bahkan google sendiri, ketika menggunakan golang, context wajib digunakan dan selalu dikirim ke setiap func yg dikirim
	*/

	// package context
	// context direpresentasikan di dalam sebuat interface Context
}
