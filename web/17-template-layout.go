package main

func main() {
	// template layout
	/*
		● Saat kita membuat halaman website, kadang ada beberapa bagian yang selalu sama, misal header
		dan footer
		● Best practice nya jika terdapat bagian yang selalu sama, disarankan untuk disimpan pada template
		yang terpisah, agar bisa digunakan di template lain
		● Go-Lang template mendukung import dari template lain
	*/

	// Import Template
	/*
		Untuk melakukan import, kita bisa menggunakan perintah berikut :
		● {{template “nama”}}, artinya kita akan meng-import template “nama” tanpa memberikan data
		apapun
		● {{template “nama” .Value}}, artinya kita akan meng-import template “nama” dengan memberikann
		data value
	*/

	// template name
	/*
		● Saat kita membuat template dari file, secara otomatis nama file nya akan menjadi nama template
		● Namun jika kita ingin mengubah nama template nya, kita juga bisa melakukan mengguakan
		perintah {{define “nama”}} TEMPLATE {{end}}, artinya kita membuat template dengan nama “nama”
	*/
}
