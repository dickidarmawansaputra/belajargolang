package main

func main() {
	// context with value
	/*
		pada saat awal membuat context, context tidak memiliki value
		kita bisa menambah sebuat value dengan data Pair(key-value) ke dalam context
		saat kita menambah value ke context, secara otomatis akan tercipta child context baru artinya original contextnya tidak akan berubah sama sekali (seperti di materi sebelumnya context itu bersifat immutable / tidak bisa diubah)
		untuk membuat menambahkan value ke context, kita bisa menggunakan func context.WithValue(parent, key, value)
	*/

	//
}
