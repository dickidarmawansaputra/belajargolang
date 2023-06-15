package main

func main() {
	// timer (time.Timer)
	/*
		timer = representasi satu kejadian
		ketike waktu timer suda expire, maka event akan dikirim ke dalam channel
		untuk membuat timer bisa menggunakan time.NewTimer(duration)
	*/

	// time.AfterFunc()
	/*
		kadang ada kebutuhan ingin menjalankan sebuah func dengan delay waktu tertentu
		kita bisa memanfaatkan timer dengan menggunakan func time.AfterFunc()
		kita tidak perlu lagi menggunakan channelnya, cukup kirimkkan func yg akan dipanggil ketika timer mengirim kejadian
	*/
}
