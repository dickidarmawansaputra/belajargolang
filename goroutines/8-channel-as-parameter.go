package main

func main() {
	// channel sebagai parameter
	/*
		dalam kenyataan pembuatan aplikasi, seringnya akan mengirim channel ke func lain via parameter
		sebelumnya kita tahu di golang by default, parameter = pass by value (bukan pass by reference), artinya value akan diduplikasi lalu dikirim ke func parameter. sehingga jika ingin mengirim data asli, bisa gunakan pointer (agar pass by reference)
		berbeda dengan channel, tidak perlu melakukan hal tersebut
	*/
}
