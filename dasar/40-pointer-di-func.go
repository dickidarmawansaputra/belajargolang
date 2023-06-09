package main

import "fmt"

type Alamat struct {
	Kota, Provinsi, Negara string
}

func changeCountryToIndonesia(alamat *Alamat) {
	alamat.Negara = "Indonesia"
}
func main() {
	// saat membuat parameter di func, secara default pass by value
	// jika ubah data dalam func data aslinya tidak pernah berubah, hal ini membuat variable jadi aman karna data tidak bisa diubah
	// namun kadang ingin mengubah data asli parameter, yaitu dengan menggunakan pointer di func
	alamat := Alamat{
		Kota:     "Pontianak",
		Provinsi: "Kalimantan Barat",
		Negara:   "",
	}
	// bisa gini ubah ke pointer
	var alamatPointer *Alamat = &alamat
	changeCountryToIndonesia(alamatPointer)

	// atau langsung gini
	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat) // tidak berubah, solusinya jika pingin berubah untuk negaranya dengan pointer di parameter dan pointer di data alamat

	// jika bikin data struct lumayan gede, gunakan pointer di parameter karna jika datanya terlalu besar setiap panggil func akan diduplikat di memory & bikin memory bengkak
}
