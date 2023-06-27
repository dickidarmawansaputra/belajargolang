package main

func main() {
	// web dinamis
	/*
		● Sampai saat ini kita hanya membahas tentang membuat response menggunakan String dan juga
		static file
		● Pada kenyataannya, saat kita membuat web, kita pasti akan membuat halaman yang dinamis, bisa
		berubah-ubah sesuai dengan data yang diakses oleh user
		● Di Go-Lang terdapat fitur HTML Template, yaitu fitur template yang bisa kita gunakan untuk
		membuat HTML yang dinamis
	*/

	// HTML template
	/*
		● Fitur HTML template terdapat di package html/template
		● Sebelum menggunakan HTML template, kita perlu terlebih dahulu membuat template nya
		● Template bisa berubah file atau string
		● Bagian dinamis pada HTML Template, adalah bagian yang menggunakan tanda {{ }}
	*/

	// Membuat Template
	/*
		● Saat membuat template dengan string, kira perlu memberi tahu nama template nya
		● Dan untuk membuat text template, cukup buat text html, dan untuk konten yang dinamis, kita bisa
		gunakan tanda {{.}}, contoh :
		● <html><body>{{.}}</body></html>
	*/

	// Template Dari File
	/*
		● Selain membuat template dari string, kita juga bisa membuat template langsung dari file
		● Hal ini mempermudah kita, karena bisa langsung membuat file html
		● Saat membuat template menggunakan file, secara otomatis nama file akan menjadi nama template
		nya, misal jika kita punya file simple.html, maka nama template nya adalah simple.html
	*/

	// Template Directory
	/*
		● Kadang biasanya kita jarang sekali menyebutkan file template satu persatu
		● Alangkah baiknya untuk template kita simpan di satu directory
		● Go-Lang template mendukung proses load template dari directory
		● Hal ini memudahkan kita, sehingga tidak perlu menyebutkan nama file nya satu per satu
	*/

	// Template dari Go-Lang Embed
	/*
		● Sejak Go-Lang 1.16, karena sudah ada Go-Lang Embed, jadi direkomendasikan menggunakan
		Go-Lang embed untuk menyimpan data template
		● Menggunakan Go-Lang embed menjadi kita tidak perlu ikut meng-copy template file lagi, karena
		sudah otomatis di embed di dalam distribution file
	*/
}
