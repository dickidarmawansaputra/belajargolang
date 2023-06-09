package main

import "fmt"

type User struct {
	Name string
}

func (user User) sayHelloUser(name string) {
	fmt.Println("Hello", name, "My Name is", user.Name)
}

func main() {
	// struct bisa digunakan sebagai parameter sama seperti tipe data lainnya
	// bisa menambahkan func dalam struct, seakan-akan struct memiliki func. padahal tidak func tetap berdiri sendiri
	var dicki User
	dicki.Name = "Dicki"
	dicki.sayHelloUser("Roy")
}
