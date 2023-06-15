package main

func main() {
	// ticker (time.Ticker)
	/*
		ticker = representasi kejadian yg berulang
		ketika waktu ticker sudah expire, maka event akan dikirim ke channel
		untuk membuat ticker, bisa menggunakan time.NewTicker(duration)
		untuk menghentikan ticker, bisa menggunakan Ticker.Stop()
	*/

	// time.Tick()
	/*
		kadang kita tidak butuh data tickernya, kita hanya butuh channelnya saja
		jika demikian, bisa menggunakan func time.Tick(duration), func ini tidak akan mengembalikan ticker, hanya mengembalikan channel timernya saja
	*/
}
