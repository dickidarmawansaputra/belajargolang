package main

func main() {
	// race codition
	// salah satu masalah yg sering terjadi di dalam concurrency / parallel programming
	// masalah dengan goroutine
	/*
		saat menggunakan goroutine, dia tidak hanya berjalan secara concurrent, tapi bisa parallel juga. karena bisa ada beberapa thread yg berjalan secara parallel
		hal ini sangat berbahaya ketika kita melakukan manipulasi data variable yg sama oleh beberapa goroutine secara bersamaan
		hal ini bisa menyebabkan msalah yg bernama RACE CONDITION
	*/
}
