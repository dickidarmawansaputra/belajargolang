package main

func main() {
	// menggagalkan unit test
	/*
		menggagalkan unit test dengan panic bukan hal yang bagus karna jika error program akan berhenti, sedangkan unit test lain selanjutnya tidak bisa dieksekusi karna program berhenti
		golang menyediakan cara untuk menggagalkan unit test dengan testing.T
		terdapat func Fail(), FailNow(), Error(), dan Fatal() jika ingin menggagalkan unit test
	*/

	// t.Fail() & t.FailNow()
	/*
		Fail() akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test. namun diakhir ketika selesai, maka unit test tersebut dianggap gagal
		FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test
	*/

	// t.Error(args...) & t.Fatal(args...)
	/*
		Error() func lebih seperti log (print) error, namun setelah lakukan log error secara otomatis akan memanggil func Fail()
		Fatal() mirip Error() func namun setelah log error memanggil func FailNow()
	*/

	// best practice pake Error() atau Fatal() agar dapat menyisipkan error message
}
