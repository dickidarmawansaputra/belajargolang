package main

func main() {
	// membuat goroutine
	/*
		untuk membuat goroutine cukup menambahkan perintah go sebelum memanggil func yg akan kita jalankan dalam goroutine (sudah built in di golang tanpa perlu import package)
		saat sebuah func dijalankan dalam goroutine, func tersebut berjalan secara asynchronous, artinya tidak akan menunggu sampai func tersebut selesai
		aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yg kita buat selesai
		(hati-hati saat gunakan goroutine, kalo tiba2 aplikasi berenti maka goroutine nggak akan sampe ditunggui goroutinenya kelar dulu)
	*/
	// cek di file test/goroutine_test.go
}
