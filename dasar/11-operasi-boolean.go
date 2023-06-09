package main

import "fmt"

func main() {
	// (&& and) (|| or) (! negasi)
	var ujian = 90
	var presensi = 80

	var lulusUjian bool = ujian >= 80
	var lulusPresensi bool = presensi >= 80

	fmt.Println(lulusUjian)
	fmt.Println(lulusPresensi)
	fmt.Println(ujian >= 80 && ujian >= 80)
}
