package main

import "fmt"

func main() {
	name := "Dicki"

	switch name {
	case "Dicki":
		fmt.Println("Hello", name)
	case "Darmawan":
		fmt.Println("Helo", name)
	default:
		fmt.Println("Hi there")
	}

	// switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama panjang")
	case false:
		fmt.Println("Nama pendek")
	}

	// switch tanpa kondisi, mirip if statement
	// bisa sebagai pengganti if statement jika kondisi sederhana
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
