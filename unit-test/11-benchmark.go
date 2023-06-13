package main

func main() {
	// golang testing package juga mendukung untuk benchmark
	// benchmark = mekanisme menghitung kecepatan performa kode aplikasi
	// benchmark di golang dilakukan secara otomatis melakukan iterasi kode yg kita panggil berkali-kali sampai waktu tertentu
	// kita tidak perlu menentukan jumlah iterasi dan lamanya, karena itu sudah diatur oleh testing.B bawaan dari testing package

	// testing.B
	/*
		testing.B = struct yg digunakan untuk melakukan benchmark
		testing.B mirip dengan testing.T, terdapat func Fail(), FailNow(), Error(), Fatal() dan lain-lain
		yg membedakan, ada beberapa attribute dan func tambahan yg digunakan untuk melakukan benchmark
		salah satunya attribute N, ini digunakan untuk melakukan total iterasi sebuah benchmark
	*/

	// cara kerja benchmark
	/*
		cara kerja benchmark di golang sederhana
		kita hanya perlu membuat perulangan sejumlah N attribute
		secara otomatis golang akan melakukan eksekusi sejumlah perulangan yang ditentukan secara otomatis, lalu mendeteksi berapa lama proses tersebut berjalan, dan disimpulkan performa benchmarknya dalam waktu
	*/
}
