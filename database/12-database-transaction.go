package main

func main() {
	// database transaction di golang
	/*
		● Secara default, semua perintah SQL yang kita kirim menggunakan Golang akan otomatis di commit,
		atau istilahnya auto commit
		● Namun kita bisa menggunakan fitur transaksi sehingga SQL yang kita kirim tidak secara otomatis di
		commit ke database
		● Untuk memulai transaksi, kita bisa menggunakan function (DB) Begin(), dimana akan menghasilkan
		struct Tx yang merupakan representasi Transaction
		● Struct Tx ini yang kita gunakan sebagai pengganti DB untuk melakukan transaksi, dimana hampir
		semua function di DB ada di Tx, seperti Exec, Query atau Prepare
		● Setelah selesai proses transaksi, kita bisa gunakan function (Tx) Commit() untuk melakukan commit
		atau Rollback()
	*/
}
