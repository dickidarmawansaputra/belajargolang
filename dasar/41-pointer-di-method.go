package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	// walaupun method akan menempel di struct, tapi sebenarnya data struct yg diakses dalam method adalah pass by value
	// sangat direkomendasikan menggunakan pointer di method sehingga tidak boros memory karna selalu di duplikasi ketika panggil method
	man := Man{"Dicki"}
	man.Married()

	fmt.Println(man.Name) // tetap dicki tanpa mr. karna tanpa pointer akan duplikasi, tambahkan pointer maka akan sesuai mr. dicki & lebih hemat memory
}
