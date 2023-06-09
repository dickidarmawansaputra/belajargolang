package main

import "fmt"

func main() {
	// beda array & slice jumlah data bisa sebanyak-banyaknya asal kata kunci berbeda/unik
	person := map[string]string{
		"name":    "Dicki",
		"address": "pontianak",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// nambahkan data
	person["title"] = "programmer"
	fmt.Println(person)

	// func map
	var book map[string]string = make(map[string]string)
	book["title"] = "tes"
	book["author"] = "Dicki"
	fmt.Println(book)

	delete(book, "author")
	fmt.Println(book)
}
