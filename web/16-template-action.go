package main

func main() {
	// template action
	/*
		● Go-Lang template mendukung perintah action, seperti percabangan, perulangan dan lain-lain
	*/

	// If Else
	/*
		● {{if .Value}} T1 {{end}}, jika Value tidak kosong, maka T1 akan dieksekusi, jika kosong, tidak ada yang
		dieksekusi
		● {{if .Value}} T1 {{else}} T2 {{end}}, jika value tidak kosong, maka T1 akan dieksekusi, jika kosong, T2
		yang akan dieksekusi
		● {{if .Value1}} T1 {{else if .Value2}} T2 {{else}} T3 {{end}}, jika Value1 tidak kosong, maka T1 akan
		dieksekusi, jika Value2 tidak kosong, maka T2 akan dieksekusi, jika tidak semuanya, maka T3 akan
		dieksekusi
	*/

	// Operator Perbandingan
	/*
		Go-Lang template juga mendukung operator perbandingan, ini cocok ketika butuh melakukan
		perbandingan number di if statement, berikut adalah operator nya :
		● eq artinya arg1 == arg2
		● ne artinya arg1 != arg2
		● lt artinya arg1 < arg2
		● le artinya arg1 <= arg2
		● gt artinya arg1 > arg2
		● ge artinya arg1 >= arg2
	*/

	// Kenapa Operatornya di Depan?
	/*
		● Hal ini dikarenakan, sebenarnya operator perbandingan tersebut adalah sebuah function
		● Jadi saat kita menggunakan {{eq First Second}}, sebenarnya dia akan memanggil function eq
		dengan parameter First dan Second : eq(First, Second)
	*/

	// range
	/*
		● Range digunakan untuk melakukan iterasi data template
		● Tidak ada perulangan biasa seperti menggunakan for di Go-Lang template
		● Yang kita bisa lakukan adalah menggunakan range untuk mengiterasi tiap data array, slice, map
		atau channel
		● {{range $index, $element := .Value}} T1 {{end}}, jika value memiliki data, maka T1 akan dieksekusi
		sebanyak element value, dan kita bisa menggunakan $index untuk mengakses index dan $element
		untuk mengakses element
		● {{range $index, $element := .Value}} T1 {{else}} T2 {{end}}, sama seperti sebelumnya, namun jika
		value tidak memiliki element apapun, maka T2 yang akan dieksekusi
	*/

	// with
	/*
		● Kadang kita sering membuat nested struct
		● Jika menggunakan template, kita bisa mengaksesnya menggunakan .Value.NestedValue
		● Di template terdapat action with, yang bisa digunakan mengubah scope dot menjadi object yang
		kita mau
		● {{with .Value}} T1 {{end}}, jika value tidak kosong, di T1 semua dot akan merefer ke value
		● {{with .Value}} T1 {{else}} T2 {{end}}, sama seperti sebelumnya, namun jika value kosong, maka T2
		yang akan dieksekusi
	*/

	// comment
	/*
		● Template juga mendukung komentar
		● Komentar secara otomatis akan hilang ketika template text di parsing
		● Untuk membuat komentar sangat sederhana, kita bisa gunakan
	*/
	// {{/* Contoh Komentar */}}
}
