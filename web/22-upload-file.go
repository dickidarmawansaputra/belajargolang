package main

func main() {
	// upload file
	/*
		● Saat membuat web, selain menerima input data berupa form dan query param, kadang kita juga
		menerima input data berupa file dari client
		● Go-Lang Web sudah memiliki fitur untuk management upload file
		● Hal ini memudahkan kita ketika butuh membuat web yang menerima input file upload
	*/

	// MultiPart
	/*
		● Saat kita ingin menerima upload file, kita perlu melakukan parsing terlebih dahulu menggunakan
		Request.ParseMultipartForm(size), atau kita bisa langsung ambil data file nya menggunakan
		Request.FormFile(name), di dalam nya secara otomatis melakukan parsing terlebih dahulu
		● Hasilnya merupakan data-data yang terdapat pada package multipart, seperti multipart.File
		sebagai representasi file nya, dan multipart.FileHeader sebagai informasi file nya
	*/
}
