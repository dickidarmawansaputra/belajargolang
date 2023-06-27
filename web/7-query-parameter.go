package main

func main() {
	// query parameter
	/*
		● Query parameter adalah salah satu fitur yang biasa kita gunakan ketika membuat web
		● Query parameter biasanya digunakan untuk mengirim data dari client ke server
		● Query parameter ditempatkan pada URL
		● Untuk menambahkan query parameter, kita bisa menggunakan ?nama=value pada URL nya
	*/

	// url.URL
	/*
		● Dalam parameter Request, terdapat attribute URL yang berisi data url.URL
		● Dari data URL ini, kita bisa mengambil data query parameter yang dikirim dari client dengan
		menggunakan method Query() yang akan mengembalikan map
	*/

	// Multiple Query Parameter
	/*
		● Dalam spesifikasi URL, kita bisa menambahkan lebih dari satu query parameter
		● Ini cocok sekali jika kita memang ingin mengirim banyak data ke server, cukup tambahkan query
		parameter lainnya
		● Untuk menambahkan query parameter, kita bisa gunakan tanda & lalu diikuti dengan query
		parameter berikutnya
	*/

	// Multiple Value Query Parameter
	/*
		● Sebenarnya URL melakukan parsing query parameter dan menyimpannya dalam
		map[string][]string
		● Artinya, dalam satu key query parameter, kita bisa memasukkan beberapa value
		● Caranya kita bisa menambahkan query parameter dengan nama yang sama, namun value berbeda,
		misal :
		● name=Eko&name=Kurniawan
	*/
}
