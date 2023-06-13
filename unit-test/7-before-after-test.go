package main

func main() {
	// biasanya dalam unit test, ketika ingin melakukan sesuatu sebelum dan setelah sebuat unit test dieksekusi
	// jika kode yg dilakukan sebelum dan sesudah selalu sama antar unit test func, maka membuat manual di unit test func adalah hal yg membosankan dan terlalu banyak kode duplikat
	// di golang ada fitur testing.M
	// fitur ini bernama main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga bisa digunakan untuk melakukan before & after di unit test

	// testing.M
	/*
		untuk mengatur eksekusi unit test, cuku membuat func bernama TestMain dengan parameter testing.M
		jika terdapat func TestMain, maka secara otomatis golang akan mengeksekusi func ini tiap kali akan menjalankan unit test di sebuah package
		dengan ini bisa mengatur before & after unit test sesuai dengan yg dimau
		ingat, func TestMain dieksekusi hanya sekali per golang package, bukan per tiap func unit test
	*/

	// walaupun ini bukan fitur before & after, tapi bisa memanfaatkan sebagai before & after unit test
}
