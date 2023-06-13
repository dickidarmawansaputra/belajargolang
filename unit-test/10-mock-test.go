package main

func main() {
	// mock = objeck yg sudah kita program dengan ekspetasi tertentu sehingga ketika dipanggil akan menghasilkan data yg sudah diprogram sebelumnya
	// mock = salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object yg memang sulit untuk di testing
	// misal membuat unit test, namun ada kode program yang harus memanggil API CALL ke third party service, dan belum tentu responsenya sesuai dengan apa yg dimau
	// pada kasus ini, cocok sekali dengan mock object

	// testify mock
	/*
		untuk membuat mock object, tidak ada fitur bawaan di golang. namun bisa menggunakan library testify
		testify mendukung pembuatan mock object, sehingga cocok untuk digunakan ketika ingin membuat mock object
		namun perlu diperhatikan, jika desain kode program jelek akan sulit untuk melakukan mocking.
		jadi pastikan melakukan desain kode program dengan baik
	*/

	// contoh kasus
	// aplikasi query ke database
	/*
		contoh kasus kali ni membuat aplikasi golang melakukan query ke database
		dimana akn membuat layer service sebagai business logic, dan layer repositoru sebagai jembatan ke database
		agar kode mudah untuk di test, disankan membuat kontrak berupa interface
	*/

}
