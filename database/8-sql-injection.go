package main

func main() {
	// sql dengan parameter
	/*
		● Saat membuat aplikasi, kita tidak mungkin akan melakukan hardcode perintah SQL di kode Golang
		kita
		● Biasanya kita akan menerima input data dari user, lalu membuat perintah SQL dari input user, dan
		mengirimnya menggunakan perintah SQL
	*/

	// sql injection
	/*
		● SQL Injection adalah sebuah teknik yang menyalahgunakan sebuah celah keamanan yang terjadi
		dalam lapisan basis data sebuah aplikasi.
		● Biasa, SQL Injection dilakukan dengan mengirim input dari user dengan perintah yang salah,
		sehingga menyebabkan hasil SQL yang kita buat menjadi tidak valid
		● SQL Injection sangat berbahaya, jika sampai kita salah membuat SQL, bisa jadi data kita tidak aman
	*/

	// solusinya
	/*
		● Jangan membuat query SQL secara manual dengan menggabungkan String secara bulat-bulat
		● Jika kita membutuhkan parameter ketika membuat SQL, kita bisa menggunakan function Execute
		atau Query dengan parameter yang akan kita bahas di chapter selanjutnya
	*/
}
