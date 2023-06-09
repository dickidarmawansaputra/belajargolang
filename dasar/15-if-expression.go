package main

import "fmt"

func main() {
	name := "Dicki"
	if name == "Dicki" {
		fmt.Println("Hello", name)
	} else if name == "Roy" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Hi there")
	}

	// tanpa short statement
	var length = len(name)

	if length > 5 {
		fmt.Println("nama panjang")
	} else {
		fmt.Println("nama pendek")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("nama panjang")
	} else {
		fmt.Println("nama pendek")
	}
}
