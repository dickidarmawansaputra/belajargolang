package main

import "fmt"

func main() {
	// variable hanya bisa menyimpan tipe data yang sama
	// name variable bersifat unik, jika var a = 1 kemudian buat lagi var a = 2 maka return value = 2
	// setiap variable yang dibuat harus digunakan jika tidak akan error
	var name string
	name = "Dicki Darmawan Saputra"
	fmt.Println(name)

	// jika ingin rubah value variable
	name = "Dicki Darmawan Saputra Dirubah"
	fmt.Println(name)

	// jika rubah value dengan tipe data berbeda akan error
	// name = true
	// fmt.Println(name)

	// jika variable langsung tanpa value, tidak perlu menyebutkan tipe datanya. secara otomatis golang akan tahu
	var nickname = "Dicki"
	var age = 17
	fmt.Println(nickname)
	fmt.Println(age)

	// kata kunci var tidak wajib, sebagai ganti dengan kata kunci :=
	// kata kunci ini cuma untuk deklarasi awal, jika ingin merubah valuenya cukup pake =
	country := "indonesia"
	fmt.Println(country)

	// multiple variable
	var (
		firstName = "Dicki"
		lastName  = "Darmawan Saputra"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
