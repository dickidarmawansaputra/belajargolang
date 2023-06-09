package main

import "fmt"

func getMyName() (firstName, middleName, lastName string) {
	firstName = "Dicki"
	middleName = "Darmawan"
	lastName = "Darmawan"

	return // sama seperti return firstName, middleName, lastName
}

func main() {
	// ketika func mengembalikan value biasa deklarasikan tipe data return valuenya
	// fitur ini bisa buat variable langsung di tipe data return functionnya
	firstName, middleName, lastName := getMyName()
	fmt.Println(firstName, middleName, lastName)
}
