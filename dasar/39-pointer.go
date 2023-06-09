package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// gunakan operator & untuk membuat pointer di variable
	// gunakan * untuk pointer misal *Address ponter ke struct misalnya
	// secara default di go semua variable passing by value bukan by reference
	// artinya mengirim sebuah variable ke func, method, atau variable lain sebenarnya adalah duplikasi valuenya
	// contohnya seperti ini
	address1 := Address{"Pontianak", "Kalimantan Barat", "Indonesia"}
	address2 := address1
	var address4 *Address = &address1
	address2.City = "Bandung"

	fmt.Println(address1) // harusnya data address1 berubah, karna di go passing by value jadi ini bukti jika sebenarnya duplikasi dari valuenya tidak seperti by reference seperti bahasa lain
	fmt.Println(address2)

	// pointer = kemampuan membuat reference ke lokasi data di memory yg sama tanpa duplikasi data yg sudah ada
	// artinya kemampuan pointer yaitu kita bisa membuat passing by reference dari sebelumnya yang passing by value
	// caranya untuk variable dengan nilai pointer ke variable yg lain dengan menambahkan operator & diikuti nama variablenya
	address3 := &address1
	address3.City = "Bengkayang"

	fmt.Println("Dengan Pointer")
	fmt.Println(address1) // akan berubah data address1
	fmt.Println(address3) // jika hover address3 akan ada tanda * tandanya pointer

	// saat mengubah variable pointer, maka yg berubah hanya variable tersebut
	// semua variable yg mengacu ke data yg sama tidak akan berubah
	// jika ingin mengubah semua variable yg mengacu ke data tersebut, bisa dengan operator *
	// contoh yg berubah hanya variable pointer
	fmt.Println(address2) // data city tetap bandung tidak berubah menjadi bengkayang karna yg berubah hanya variable pointer

	*address3 = Address{"Sanggau Ledo", "Kalimantan Barat", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address3)
	fmt.Println(address4)

	alamat1 := Address{"Bange", "Kalimantan Barat", "Indonesia"}
	var alamat2 *Address = &alamat1
	var alamat3 *Address = &alamat1
	alamat2.City = "Kandasan"

	*alamat2 = Address{"Balikpapan", "Kalimantan Barat", "Indonesia"}

	fmt.Println(alamat1)
	fmt.Println(alamat2)
	fmt.Println(alamat3)

	// func new untuk membuat pointer dengan data kosong = tidak ada data awal
	addressNew := new(Address)
	addressNew.City = "Jakarta"
	fmt.Println(addressNew)
}
