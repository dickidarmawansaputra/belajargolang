package main

func main() {
	// context with cancel
	/*
		selain menambahkan value ke context, kita juga bisa menambahkan sinyal cancel ke context
		kapan sinyal cancel diperlukan dalam context?
		biasanya ketika kita butuh menjalankan proses lain, dan kita ingin bisa memberi sinyal cancel ke proses tersebut
		biasanya proses ini berupa goroutine yg berbeda, sehingga dengan mudah jika kita ingin membatalkan eksekusi goroutine, kita bisa mengirim sinyal cancel ke contextnya
		namun ingat, goroutine yg menggunakan context tetap harus melakukan pengecekan terhadap contextnya, jika tidak maka tidak ada gunanya (bikin parameter context tapi tidak pernah di pake di goroutine, jadi percuma kalo tidak dipake goroutine tidak akan berhenti)
		untuk membuat context dengan cancel signal, kita bisa menggunakan func context.WithCancel(parent)
	*/
}
