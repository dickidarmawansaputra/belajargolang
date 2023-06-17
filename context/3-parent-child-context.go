package main

func main() {
	// parent dan child context
	/*
		context menganut konsep parent dan child (adanya pewarisan)
		artinya, saat membuat context, bisa membuat child context dari context yg sudah ada
		parent context bisa memiliki banyak child, namun child hanya bisa memiliki satu parent context
		konsep ini mirip dengan pewarisan di pemrograman berorientasi object
	*/

	// hubungan antara parent dan child context
	/*
		parent dan child context akan selalu terhubung
		saat nanti melakukan misal pembatalan context A, maka semua child dan sub child dari context A akan ikut dibatalkan
		namun jika misal kita membatalkan context B, hanya context B dan semua child dan sub childnya yg dibatalkan, parent context B tidak akan ikut dibatalkan
		begitu juga nanti saat kita menyisipkan data ke dalam context A, semua child dan sub childnya bisa mendapatkan data tersebut
		namun jika kita menyisipkan data di context B, hanya context B dan semua child dan sub childnya yang mendapat data, parent context B tidak akan mendapat data
	*/

	// immutable
	/*
		context = object yg immutable, artinya setelah context dibuat, dia tidak bisa diubah lagi (cuma bisa baca datanya, tidak bisa dirubah datanya jadi aman)
		ketika kita menambahkan value ke dalam context, atau menambahkan pengaturan timeout dan yg lainnya, secara otomatis akan membentuk child context baru, bukan merubah context tesebut
	*/

	// cara membuat child context
	// cara membuat child context ada banyak caranya, next selanjutnya
}
