package main

func main() {
	// query sql
	/*
		● Untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggunakan perintah Exec, namun
		jika kita membutuhkan result, seperti SELECT SQL, kita bisa menggunakan function yang berbeda
		● Function untuk melakukan query ke database, bisa menggunakan function (DB)
		QueryContext(context, sql, params)
	*/

	// rows
	/*
		● Hasil Query function adalah sebuah data structs sql.Rows
		● Rows digunakan untuk melakukan iterasi terhadap hasil dari query (mirip cursor)
		● Kita bisa menggunakan function (Rows) Next() (boolean) untuk melakukan iterasi terhadap data
		hasil query, jika return data false, artinya sudah tidak ada data lagi didalam result
		● Untuk membaca tiap data, kita bisa menggunakan (Rows) Scan(columns...)
		● Dan jangan lupa, setelah menggunakan Rows, jangan lupa untuk menutupnya menggunakan
		(Rows) Close()
	*/
}
