package main

import "fmt"

// penamaan biasanya dari huruf besar
type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	// stuct = sebuah template data yg digunakan untuk menggabungkan nol / lebih tipe data dalam satu kesatuan
	// biasanya bisa pake array, cuma problemnya array datanya harus setipe data dan akses datanya pake index
	// jika map key nya harus sama tipe datanya
	// struct biasanya representasi data, disimpan dalam field
	// sederhananya kumpulan dari field
	// kalo di OOP mau buat data customer berarti class Customer mirip seperti itu
	// stuct = template data / prototype data & tidak bisa digunakan langsung
	// jika ingin bikin data bagusnya pake struct lebih terstruktur dibanding pake map / array
	var dicki Customer
	dicki.Name = "Dicki"
	dicki.Address = "Pontianak"
	dicki.Age = 17

	fmt.Println(dicki)
	fmt.Println(dicki.Name)

	// Struct Literals
	// ada banyak cara membuat struct
	user := Customer{
		Name:    "Dicki",
		Address: "Pontianak",
		Age:     17,
	}
	fmt.Println(user)

	// cara ini harus sesuai urutan, jika tidak maka akan error
	// tidak disarankan pake cara ini, jika ada perubahan di struct maka akan error juga
	fmt.Println(Customer{"Dicki", "Pontinak", 17})

}
