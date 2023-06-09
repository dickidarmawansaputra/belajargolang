package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	// bisanya di bahasa lain object belum diinisialisakan maka otomatis nilanya null / nil
	// beda di go, ketika buat variable dengan tipe data tertentu, maka otomatis dibuatkan default valuenya
	// nil hanya bisa digunakan di beberapa tipe data seperti interface, func, map, slice, pointer, dan channel
	var tanpaNil map[string]string
	if tanpaNil["name"] == "" {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(tanpaNil)
	}

	var denganNil map[string]string = NewMap("Dicki")
	if denganNil == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(denganNil)
	}
}
