package main

func main() {
	// assertion
	/*
		melakukan pengecekan unit test secara manual dengan if else sangat menyebalkan
		oleh karna itu sangat disarankan dengan assertion untuk melakukan pengecekan
		sayangnya golang tidak menyediakan package untuk assertion sehingga perlu menambahkan library untuk assertion ini
		salah satu library paling popular di golang adalah Testify
	*/

	// assert vs require
	/*
		saat menggunakan assert, jika gagal maka assert akan memanggail Fail(), artinya eksekusi unit test tetap dilanjutkan
		jika require, saat pengecekan gagal akan memanggil FailNow(), artinya eksekusi unit test tidak akan dilanjutkan
	*/
}
