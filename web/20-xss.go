package main

func main() {
	// XSS (cross site scripting)
	/*
		● XSS adalah salah satu security issue yang biasa terjadi ketika membuat web
		● XSS adalah celah keamanan, dimana orang bisa secara sengaja memasukkan parameter yang
		mengandung JavaScript agar dirender oleh halaman website kita
		● Biasanya tujuan dari XSS adalah mencuri cookie browser pengguna yang sedang mengakses
		website kita
		● XSS bisa menyebabkan account pengguna kita diambil alih jika tidak ditangani dengan baik
	*/

	// Auto Escape
	/*
		● Berbeda dengan bahasa pemrograman lain seperti PHP, pada Go-Lang template, masalah XSS
		sudah diatasi secara otomatis
		● Go-Lang template memiliki fitur Auto Escape, dimana dia bisa mendeteksi data yang perlu
		ditampilkan di template, jika mengandung tag-tag html atau script, secara otomatis akan di escape
		● Semua function escape bisa diliat disini :
		● https://github.com/golang/go/blob/master/src/html/template/escape.go
		● https://golang.org/pkg/html/template/#hdr-Contexts
	*/

	// Mematikan Auto Escape
	/*
		● Jika kita mau, auto escape juga bisa kita matikan
		● Namun, kita perlu memberi tahu template secara eksplisit ketika kita menambahkan template data
		● Kita bisa menggunakan data
		● template.HTML , jika ini adalah data html
		● template.CSS, jika ini adalah data css
		● template.JS, jika ini adalah data javascript
	*/

	// Masalah XSS (Cross Site Scripting)
	/*
		● Saat kita mematikan fitur auto escape, bisa dipastikan masalah XSS akan mengintai kita
		● Jadi pastikan kita benar-benar percaya terhadap sumber data yang kita matikan auto escape nya
	*/
}
