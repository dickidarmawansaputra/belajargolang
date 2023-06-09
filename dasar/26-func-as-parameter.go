package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

// tanpa type declaration
// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func toUppercase(name string) string {
	if len(name) <= 3 {
		return strings.ToUpper(name)
	} else {
		return strings.ToLower(name)
	}

}

func main() {
	// jika sebuah func terlalu panjang & ribet nuliskan parameternya
	// bisa dengan type declarations untuk membuat alias func agar mempermudah func sebagai parameter
	sayHelloWithFilter("Dicki", spamFilter)
	sayHelloWithFilter("Dicki", toUppercase)

}
