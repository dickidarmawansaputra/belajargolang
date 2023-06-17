package main

func main() {
	// context with deadline
	/*
		selain menggunakan timeout untuk melakukan cancel secara otomatis, juga bisa menggunakan deadline
		pengaturan deadline sedikit berbeda dengan timeout, jika timeout kita beri waktu dari sekarang, kalo deadline ditentukan kapan waktu timeoutnya, misal jam 12 siang hari ini
		untuk membuat context dengan cancel signal secara otomatis menggunakan deadline, bisa menggunakan func context.WithDeadline(parent, time)
	*/
}
