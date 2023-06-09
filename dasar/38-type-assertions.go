package main

import "fmt"

func random() interface{} {
	return "OK"
}
func main() {
	// type assertions = kemampuan merubah tipe data jadi tipe data yang diinginkan
	// fitur ini sering gunakan ketika bertemu interface kosong
	// hati-hati saat gunakan fitur ini & yakin dari tipe data yg digunakan
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	// resultInt := result.(int) // panic
	// fmt.Println(resultInt) // panic

	// jika ingin aman gunakan type assertion dengan switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown type")
	}

}
