package main

func main() {
	// RWMutex (read write mutex)
	/*
		kadang ada kasus dimana ingin melakukan locking tidak hanya pada proses mengubah data, tapi membaca data
		sebenarnya bisa menggunakan mutex saja, namun masalahnya nanti akan rebutan antara proses membaca dan mengubah
		di golang telah disediakan struct RWMutex (read write mutex) untuk menangani hal ini, dimana mutex jenis ini memiliki dua lock,, lock untuk read dan lock untuk write
	*/
}
