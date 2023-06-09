package main

import "fmt"

func getFullName() (string, string) {
	return "Dicki", "Darmawan Saputra"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// untuk ignore variable yang tidak dibutuhkan gunakan kata kunci _
	first, _ := getFullName()
	fmt.Println(first)
}
