package main

func main() {
	// buffered channel
	/*
		seperti yg dijelaskan sebelumnya, secara default channel itu hanya bisa menerima 1 data
		artinya jika menambah data ke 2, maka akan diminta menunggu sampai data ke 1 ada yg mengambil
		kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita menggunakan channel, maka otomatis pengirim akan ikut lambat juga
		untuk itu ada buffered channel, yaitu buffer yang bisa digunakan untuk menampung data antrian di channel
	*/

	// buffer capacity
	/*
		kita bebas memasukan berapa jumlah kapasitas antrian dalam buffer (tetapi harus bijak, jangan kegedean juga)
		jika kita set misal 5, artinya bisa menerima 5 data buffer
		jika kita mengirim data ke 6, maka kita diminta untuk menunggu sampai buffer ada yang kosong
		ini cocok sekali ketika memang goroutine yg menerima data lebih lambat dari yg mengirim data
	*/
}
