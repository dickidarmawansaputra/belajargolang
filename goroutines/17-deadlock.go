package main

func main() {
	// deadlock
	// masalah yg terjadi juga pada concurrent & parallel selain race condition yaitu deadlock
	// bisanya masalah ini terjadi karna salah implementasi mutex atau locking
	/*
		hati-hati saat membuat aplikasi parallel atau concurrent, masalah yg sering dihadapi adalah deadlock
		deadlock = keadaan dimana sebuat proses goroutine saling menunggu lock sehingga tidak ada satupun goroutine yang bisa jalan
		berikut simulasi proses deadlock di file test/deadlock_test.go
	*/
}
