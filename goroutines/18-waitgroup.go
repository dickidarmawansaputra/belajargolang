package main

func main() {
	// WaitGroup
	/*
		waitgroup = fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan
		hal ini kadang perlu dilakukan, misal ingin menjalankan beberapa proses menggunakan goroutine, tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai
		kasus seperti ini bisa menggunakan waitgroup
		untuk menandai bahwa ada proses goroutine, kita bisa menggunakan method Add(int), setekah proses goroutine selesai, kiga bisa gunakan method Done()
		untuk menunggu semua proses selesai, kita bisa menggunakan method Wait()
	*/
}
