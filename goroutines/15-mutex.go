package main

func main() {
	// sync.Mutex
	// mutex (mutual exclusion)
	/*
		untuk mengatasi race condition tersebut, di golang terdapat sebuat struct bernama sync.Mutex
		mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika melakukan locking terhadap mutex, maka tidak ada yg bisa melakukan locking lagi sampai kita melakukan unlock
		dengan demikian, jika ada beberapa goroutine melakukan lock terhadap mutex, maka hanya 1 goroutine yang diperbolehkan, setelah goroutine melakukan unlock baru goroutine selanjutnya diperbolehkan melakukan lock lagi
		ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya dihadapi
	*/
}
