package main

import (
	"fmt"
)

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	// letakan diatas agar jika error tetap dieksekusi
	defer logging()
	fmt.Println("Jalankan aplikasi")
	result := 10 / value
	fmt.Println(result)
}

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}
func endAppWithRecover() {
	message := recover()
	if message != nil {
		fmt.Println(message)
	} else {
		fmt.Println("Aplikasi selesai")
	}
}

func runAppWithRecover(error bool) {
	defer endAppWithRecover()
	if error {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	// defer func = func yang bisa dijadwalkan untuk dieksekusi setelah sebuat func selesai di eksekusi
	// akan selalu di eksekusi walaupun terjadi error di func yang dieksekusi
	// harusnya error karna dibagi 0 dan tidak bisa panggil logging()
	// karna pake defer makan logging tetap akan terpanggil
	runApplication(0)

	// panic func = func yg digunakan untuk menghentikan program
	// biasa dipanggil ketika terjadi error pada program berjalan
	// namun defer func tetap akan dipanggil
	runApp(true)

	// recover = func yang bisa digunakan untuk menangkap data panic
	// dengan recover proses panic akan terhenti sehingga program tetap jalan
	// agar ketika terjadi error recover bisa menangkap panic message, letakan recover di defer func
	runAppWithRecover(true)
}
