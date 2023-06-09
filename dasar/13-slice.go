package main

import "fmt"

func main() {
	// tipe data slice = potongan dari array
	// bedanya ukuran slice bisa berubah
	// menyimpan data dalam array
	// memiliki 3 data: pointer, length, capacity
	// gunakan ... jika belum tahu kapasitas arraynya
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)

	// jika data slice diubah maka data array ikut berubah
	// ini bukan dirubah di index 0 pada array tapi index 0 pada data slice = index ke 4
	// begitupun sebaliknya jika data array dirubah maka data slice juga berubah
	slice1[0] = "berubah"
	months[5] = "berubah juga"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	// func di slice
	// jika kapasitas array penuh dan append data maka akan buat array baru
	// jika buat array baru dan rubah data slicenya makan yang berubah di array baru
	// data array lama tidak berubah
	var slice3 = append(slice2, "Dicki")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"
	fmt.Println(slice3)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Dicki"
	newSlice[1] = "Darmawan Saputra"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))

	// hati-hati saat buat array
	iniArray := [...]int{1, 2, 3}
	iniArray2 := [3]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}

	fmt.Println(iniArray)
	fmt.Println(iniArray2)
	fmt.Println(iniSlice)

}
