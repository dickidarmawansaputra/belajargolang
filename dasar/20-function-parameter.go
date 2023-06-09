package main

import "fmt"

func sayHi(firstName string, lastName string) {
	fmt.Println("Hi", firstName, lastName)
}

func main() {
	sayHi("Dicki", "Darmawan Saputra")
}
