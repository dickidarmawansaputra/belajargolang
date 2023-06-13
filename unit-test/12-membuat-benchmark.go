package main

func main() {
	// benchmark func
	/*
		mirip seperti unit test untuk di golang sudah ditentukan nama funcnya
		harus diawali dengan kata Benchmark, misal BenchmarkHelloWorld
		selain itu, harus memiliki parameter (b *testing.B)
		dan tidak boleh mengembalikan return value
		untuk nama file benchmark sama seperti unit test, diakhiri dengan _test. misal hello_world_test.go
	*/

	// menjalankan benchmark
	/*
		menjalankan seluruh benchmark di module: "go test -v -bench=."
		menjalankan benchmark tanpa unit test: "go test -v -run=NotMathUnitTest (artinya run namafunc unit test yg tidak ada) -bench=."

	*/
}
