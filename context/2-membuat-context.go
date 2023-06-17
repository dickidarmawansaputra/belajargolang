package main

func main() {
	// membuat context
	/*
		karena context adalah sebuah interface, untuk membuat context butuh sebuah struct yg sesuai dengan kontrak interface context
		namun kita tidak perlu membuat secara manual
		di golang package context terdapat func yg bisa digunakan untuk membuat context
	*/

	// func membuat context
	/*
		context.Background() = membuat context kosong, tidak pernah dibatalkan, tidak pernah timeout, dan tidak memiliki value apapun. biasanya digunakan di main func atau dalam test, atau dalam awal proses request terjadi
		context.TODO() = membuat context kosong seperti Backgroud(), namun biasanya menggunakan ini ketika belum jelas context apa yg ingin digunakan
	*/
}
