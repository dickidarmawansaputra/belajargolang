package main

func main() {
	// response code
	/*
		● Dalam HTTP, terdapat yang namanya response code
		● Response code merupakan representasi kode response
		● Dari response code ini kita bisa melihat apakah sebuah request yang kita kirim itu sukses diproses
		oleh server atau gagal
		● Ada banyak sekali response code yang bisa kita gunakan saat membuat web
		● https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
	*/

	// mengubah response code
	/*
		● Secara default, jika kita tidak menyebutkan response code, maka response code nya adalah 200 OK
		● Jika kita ingin mengubahnya, kita bisa menggunakan function ResponseWriter.WriteHeader(int)
		● Semua data status code juga sudah disediakan di Go-Lang, jadi kita ingin, kita bisa gunakan variable
		yang sudah disediakan : https://github.com/golang/go/blob/master/src/net/http/status.go
	*/
}
