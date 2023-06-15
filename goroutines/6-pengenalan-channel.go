package main

func main() {
	// pengenalan channel
	// digunakan jika butuh return data dari goroutine
	// pengirim & penerima juga goroutine
	/*
		channel = tempat komunikasi secara synchronous yg bisa dilakukan oleh goroutine
		di channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yg berbeda
		saat melakukan pengiriman data ke channel, goroutine akan ter-block, sampai ada yg menerima data tersebut
		maka dari itu, channel disebut sebagai alat komunikasi synchronous (blocking)
		channel cocok sekali sebagai alternatif seperti mekanisme async await yg terdapat di beberapa bahasa pemrograman lain
	*/

	// karakteristik channel
	/*
		secara default channel hanya bisa menampung satu data, jika ingin menambahkan data lagi harus menunggu data yg ada di channel diambil
		channel hanya bisa menerima satu jenis data (sebenarnya bisa dengan gunakan interface kosong, cuma yg mengambil data dari channel harus melakukan conversi tipe data manual dengan tipe data yg dituju)
		channel bisa diambil dari lebih satu goroutine
		channel harus di close jika tidak digunakan, atau bisa menyebabkan memory leak
	*/
	// pengirim & penerima sebenarnya bisa banyak, cuma harus satu-satu
}
