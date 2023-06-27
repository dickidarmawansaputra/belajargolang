package main

func main() {
	// ServeMux
	/*
		● Saat membuat web, kita biasanya ingin membuat banyak sekali endpoint URL
		● HandlerFunc sayangnya tidak mendukung itu
		● Alternative implementasi dari Handler adalah ServeMux
		● ServeMux adalah implementasi Handler yang bisa mendukung multiple endpoint
	*/

	// URL pattern
	/*
		● URL Pattern dalam ServeMux sederhana, kita tinggal menambahkan string yang ingin kita gunakan
		sebagai endpoint, tanpa perlu memasukkan domain web kita
		● Jika URL Pattern dalam ServeMux kita tambahkan di akhirnya dengan garis miring, artinya semua
		url tersebut akan menerima path dengan awalan tersebut, misal /images/ artinya akan dieksekusi
		jika endpoint nya /images/, /images/contoh, /images/contoh/lagi
		● Namun jika terdapat URL Pattern yang lebih panjang, maka akan diprioritaskan yang lebih panjang,
		misal jika terdapat URL /images/ dan /images/thumbnails/, maka jika mengakses
		/images/thumbnails/ akan mengakses /images/thumbnails/, bukan /images
	*/
}
