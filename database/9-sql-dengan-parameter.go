package main

func main() {
	// sql dengan parameter
	/*
		● Sekarang kita sudah tahu bahaya nya SQL Injection jika menggabungkan string ketika membuat
		query
		● Jika ada kebutuhan seperti itu, sebenarnya function Exec dan Query memiliki parameter tambahan
		yang bisa kita gunakan untuk mensubtitusi parameter dari function tersebut ke SQL query yang
		kita buat.
		● Untuk menandai sebuah SQL membutuhkan parameter, kita bisa gunakan karakter ? (tanda tanya)
	*/

	// contoh sql
	// gunakan tanda "?"
	/*
		● SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1
		● INSERT INTO user(username, password) VALUES (?, ?)
		● Dan lain-lain
	*/
}
