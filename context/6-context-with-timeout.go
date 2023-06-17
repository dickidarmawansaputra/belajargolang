package main

func main() {
	// context with timeout
	/*
		selain menambahkan value ke context, dan juga sinyal cancel, kita juga bisa menambahkan sinyal cancel ke context secara otomatis dengan menggunakan pengaturan timeout
		dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel secara manual, cancel akan otomatis di eksekusi jika waktu timeout sudah terlewati
		penggunaan context dengan timeout sangat cocok ketika misal kita melakukan query ke database atau http api, namun ingin menentukan batas maksimal timeoutnya
		untuk membuat context dengan cancel signal secara otomatis menggunakan timeout, bisa dengan func context.WithTimeout(parent, duration)
	*/
}
