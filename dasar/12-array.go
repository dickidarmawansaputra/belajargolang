package main

import "fmt"

func main() {
	// tipe data berisikan data dengan tipe yang sama
	// perlu menentukan jumlah data yang ditampung oleh array
	// daya tampung tidak bisa bertambah setelah array dibuat
	var names [3]string
	names[0] = "Dicki"
	names[1] = "Darmawan"
	names[2] = "Saputra"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat array langsung
	var values = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(values)

	// example func di array
	fmt.Println(len(values))
}
