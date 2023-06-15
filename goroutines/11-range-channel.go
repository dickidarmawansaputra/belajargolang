package main

func main() {
	// range channel
	/*
		kadang-kadang ada kasus sebuah channel dikirim data secara terus menerus oleh pengirim
		dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
		salah satu yg bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel
		ketika sebuah channel di close(), maka secara otomatis perulagan tersebut akan berhenti
		ini lebih sederhana dari pada kita melakukan pengecekan channel secara manual
	*/
}
