package main

func main() {
	// GOMAXPROCS
	/*
		sebelumnya diawal sudah bahas goroutine itu sebenarnya berjalan di dalam thread
		pertanyaannya, seberapa banyak thread yg ada di golang ketika aplikasi berjalan
		untuk mengetahui berapa jumlah thread, bisa menggunakan GOMAXPROCS, yaitu sebuah func di package runtime yg bisa digunakan untuk mengubah jumlah thread atau mengambil jumlah thread
		secara default, jumlah thread di golang sebanyak jumlah CPU di komputer
		kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan func runtime.NumCpu()
	*/
}
