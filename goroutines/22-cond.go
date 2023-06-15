package main

func main() {
	// Cond (sync.Cond)
	/*
		cond = implementasi locking berbasis kondisi
		cond membutuhkan locker (bisa menggunakan mutex atau RWMutex) untuk implementasi lockingnya, namun berbeda dengan locker biasanya, di cond terdapat func Wait() untuk menunggu apakah perlu menunggu atau tidak
		func Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu lagi, sedangkan func Broadcast() digunakan untuk memberitahu semua goroutine agar tidak perlu menunggu lagi
		untuk membuat cond, bisa menggunakan func sync.NewCond(Locker)
	*/
}
