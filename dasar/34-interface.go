package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	// interface = tipe data abstract, tidak memiliki implementasi
	// berisikan definisi-definisi method
	// bisa digunakan sebagai contract
	// sama seperti struct bisa digunakan sebagai parameter
	// setiap tipe data yg sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface
	var dicki Person
	dicki.Name = "Dicki"
	SayHello(dicki)

	cat := Animal{
		Name: "Meow",
	}
	SayHello(cat)
}
