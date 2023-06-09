package main

func main() {
	// constant = variable yang nilainya tidak bisa dirubah setelah pertama kali diberi nilai
	// wajib deklarasi value
	// berbeda dari variable jika tidak digunakan maka tidak error
	const firstName string = "Dicki"
	const lastName = "Darmawan Saputra"

	// multiple constant
	const (
		x = "x"
		y = "y"
	)

	// error jika dirubah valuenya
	// firstName = "A"
	// lastName = "B"
}
