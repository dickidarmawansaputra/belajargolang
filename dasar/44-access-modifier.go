package main

import (
	"dasar/dasar/helper"
	"fmt"
)

func main() {
	// untuk menentukan access modifier di go cukup dengan nama func atau variable
	// jika nama diawali dengan huruf besar, maka bisa diakses dari package lain
	// jika diawali dengan huruf kecil tidak bisa diakses dari package lain
	// begitu juga dengan variable
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) tidak bisa diakses dari luar karna diawali huruf kecil
	fmt.Println(helper.DiAwaliHurufBesar())
	// fmt.Println(helper.diAwaliHurufKecil()) tidak bisa diakses dari luar karna diawali huruf kecil
}
