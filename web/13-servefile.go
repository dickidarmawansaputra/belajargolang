package main

func main() {
	// ServeFile
	/*
		● Kadang ada kasus misal kita hanya ingin menggunakan static file sesuai dengan yang kita inginkan
		● Hal ini bisa dilakukan menggunakan function http.ServeFile()
		● Dengan menggunakan function ini, kita bisa menentukan file mana yang ingin kita tulis ke http
		response
	*/

	// Go-Lang Embed
	/*
		● Parameter function http.ServeFile hanya berisi string file name, sehingga tidak bisa menggunakan
		Go-Lang Embed
		● Namun bukan berarti kita tidak bisa menggunakan Go-Lang embed, karena jika untuk melakukan
		load file, kita hanya butuh menggunakan package fmt dan ResponseWriter saja
	*/
}
