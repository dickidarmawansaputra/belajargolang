package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}
func main() {
	// go bukan bahasa OOP
	// biasanya jika OOP, ada satu data parent puncak dianggap sebagai semua implementasi data dari bahasa tersebut (super class)
	// di java ada java.lang.Object
	// untuk menangani kasus ini, di go dengan interface kosong
	// interface kosong = interface tidak memiliki deklarasi method satupun, yang membuat secara otomatis semua tipe data akan menjadi implementasinya
	// contoh interface kosong
	// fmt.Prinln() pacic() recover() etc.
	kosong := Ups(2)
	var data interface{} = Ups(20)
	fmt.Println(kosong)
	fmt.Println(data)
}
