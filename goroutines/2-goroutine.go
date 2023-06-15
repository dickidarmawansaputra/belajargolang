package main

func main() {
	// pengenalan goroutine
	/*
		goroutine = sebuah thread ringan yg di kelola go runtime (sebenarnya bukan thread, dan agak susah dianalogikan jadi kebanyakan orang bilang goroutine = thread ringan padahal goroutine running dalam sebuah thread)
		ukuran goroutine sangan kecil sekitar 2kb dibangkan dengan thread biasa bisa sampe 1mb
		namun tidak seperti thread yg berjalan parallel, goroutine berjalan secara concurent
		makanya golang terkenal dengan konsumsi resource yang jauh lebih kecil dibanding yg lain
	*/

	// cara kerja goroutine
	/*
		sebenarnya, goroutine dijalankan oleh go scheduler dalam thread, dimana jumlah threadnya sebanyak GOMAXPROCS (biasanya sejumlah core CPU)
		jadi sebenarnya tidak bisa dibilang goroutine itu pengganti thread, karna goroutine sendiri berjalan di atas thread
		namun yg mempermudah, kita tidak perlu lakukan manajemen thread secara manual dan semua itu sudah diatur oleh go scheduler
	*/

	// cara kerja go scheduler
	/*
		G: goroutine
		M: thread (machine)
		P: processor
	*/
}
