package main

func main() {
	// select channel
	/*
		kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine
		lalu kita ingin mendapatkan data dari semua channel tersebut
		untuk melakukan hal tersebut, kita bisa menggunakan select channel di golang (jika menggunakan for range akan ribet karna itu digunakan untuk satu channel)
		dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data datang secara bersamaan di beberapa channel, maka akan dipilih secara random
	*/
}
