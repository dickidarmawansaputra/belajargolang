package main

func main() {
	// default channel
	/*
		apa yg terjadi jika melakukan select channel terhadap channel yg ternyata tidak ada datanya?
		maka akan menunggu sampai data ada (maka akan terjadi error deadlock)
		kadang mungkin ingin melakukan sesuatu jika misal semua channel tidak ada datanya ketika melakukan select channel
		dalam select, kita bisa menambahkan default dimana ini akan dieksekusi jika memang di semua channel yg kita select tidak ada datanya
	*/
}
