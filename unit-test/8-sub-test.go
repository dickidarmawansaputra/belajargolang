package main

func main() {
	// golang mendukung fitur pembuatan func unit test di dalam func unit test
	// fitur ini sedikit aneh & jarang sekali dimiliki di unit test di bahasa pemrograman lain
	// untuk membuat sub test, bisa menggunakan Run()

	// menjalankan hanya sub test
	// go test -run TestNamaFunc/NamaSubFunc
	// atau menjalankan semua sub test di semua func
	// go test -run /NamaSubFunc
}
