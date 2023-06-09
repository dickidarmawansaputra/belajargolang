package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}
func main() {
	// func first class citizen di golang (tidak ada anak tiri)
	// bisa dibuat sebagai tipe data, disimpan variable, dan independen
	goodbye := getGoodBye
	fmt.Println(goodbye("Dicki"))
}
