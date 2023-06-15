package main

func main() {
	// Map (sync.Map)
	/*
		golang memiliki struct bernama sync.Map
		map ini mirip golang map, namun bedanya map ini aman untuk menggunakan concurrent dengan goroutine
		ada beberapa func yg bisa digunakan di map ini:
		Store(key, value) untuk menyimpan data ke map
		Load(key) untuk mengambil data dari map menggunakan key
		Delete(key) untuk menghapus data di map menggunakan key
		Range(function(key, value)) digunakan untuk melakukan iterasi seluruh data di map
	*/
}
