package main

func main() {
	// Pool (sync.Pool)
	/*
		pool = implementasi design pattern bernama object pool pattern (dikhususkan untuk concurrency & parallel, bukan pool biasa)
		sederhananya, design pattern pool ini digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya, kita bisa mengambil dari pool, dan setelah selesai menggunakan datanya kita bisa menyimpan kembali ke pool
		implementasi pool di golang ini sudah aman dari problem race condition
	*/
	// implementasinya biasa sering digunakan untuk memanage koneksi ke database
}
