package main

func main() {
	// membuat channel
	/*
		channel di golang direpresentasikan dengan tipe data chan
		untuk membuat channel, bisa menggunakan make() mirip ketika membuat map
		namun saat pembuatan channel, harus tentukan tipe data apa yg bisa dimasukan ke dalam channel tersebut
	*/

	// lihat di test/channel_test.go

	// mengirim dan menerima data dari channel
	/*
		untuk mengirim data bisa gunakan kode: "channel <-data"
		sedangkan untuk menerima data bisa gunakan kode: "data <- channel"
		jika selesai, jangan lupa untuk menutup channel dengan menggunakan func close()
		gunakan defer close() walaupun terjadi error jika selesai dieksekusi otomatis di close
	*/
}
