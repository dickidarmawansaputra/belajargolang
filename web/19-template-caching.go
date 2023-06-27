package main

func main() {
	// template caching
	/*
		● Kode-kode diatas yang sudah kita praktekan sebenarnya tidak efisien
		● Hal ini dikarenakan, setiap Handler dipanggil, kita selalu melakukan parsing ulang template nya
		● Idealnya template hanya melakukan parsing satu kali diawal ketika aplikasinya berjalan
		● Selanjutnya data template akan di caching (disimpan di memory), sehingga kita tidak perlu
		melakukan parsing lagi
		● Hal ini akan membuat web kita semakin cepat
	*/
}
