package main

func main() {
	// parallel programming
	/*
		saat ini era multicore, dan jarang sekali prosesor sekarang single core
		semakin canggih hardware, maka software juga akan mengikuti dimana sekarang bisa dengan mudah membuat proses parallel di aplikasi
		parallel programming sederhananya = memecahkan suatu masalah dengan cara membagi menjadi lebih kecil dan dijalankan secara bersamaan pada waktu yang bersamaan juga
	*/

	// contoh parallel programming
	/*
		menjalankan beberapa aplikasi sekaligus di sistem operasi kita (office, browser, editor, dll)
		beberapa koki menyiakan makanan di restoran, dimana tiap koki membuat makanan masing-masing
		antrian di bank, dimana tiap teller melayani nasabahnya masing-masing
	*/

	// process vs thread
	/*
		process = sebuah eksekusi program vs thread = segmen dari proses
		process = konsumsi memory besar vs thread = konsumsi memory kecil
		process = saling terisolasi dengan process lain vs thread = bisa saling terhubung jika dalam process yg sama
		process = lama untuk dijalankan dimatikan vs thread = cepat untuk dijalankan dimatikan
	*/

	// parallel vs concurrency
	/*
		berbeda dengan parallel (menjalankan beberapa pekerjaan secara bersamaan), concurrency = menjalankan beberapa pekerjaan secara bergantian (tidak diwaktu yg sama)
		dalam parallel biasanya membutuhkan banyak thread, sedangkan concurrency hanya butuh sedikit thread
		golang secara default concurrency, tapi karna sekaran running aplikasi secara multicore jadinya campuran concurrency dan parallel
	*/

	// contoh concurrency
	/*
		saat makan di cafe, kita bisa makan, lalu ngobrol, lalu minum, makan lagi, dan seterusnya. tetapi kita tidak bisa pada saat bersamaan minum sambil makan, dan lainnya
	*/

	// CPU-bound
	/*
		banyak algoritma dibuat hanya membutuhkan CPU untuk menjalankannya. algoritma jenis ini sangan bergantung kecepatan CPU
		contoh paling populer = machine learning, oleh karna itu sekarang banyak teknologi machine learning yg banyak gunakan GPU karna memiliki core yg lebih banyak dibanding CPU
		jenis algoritma seperti ini tidak ada benefitnya gunakan concurrency programming, namun bisa dibantu dengan implementasi parallel programming
	*/

	// I/O-bound
	/*
		I/O-bound = kebalikan dari CPU-bound, dimana biasanya algoritma / aplikasi tergantung dengan kecepatan input output devices yg digunakan
		contoh aplikasi seperti membaca data dari file, database dll
		kebanyakan aplikasi saat ini, biasanya kita akan membuat aplikasi jenis seperti ini
		aplikasi jenis I/O-bound, walaupun bisa terbantu dengan implementasi parallel programming, tapi benefitnya akan lebih baik jika menggunakan concurrency programming
		bayangkan kita membaca data dari database, dan thread harus menunggu 1 detik untuk mendapatkan balasan dari database, padahal 1 detik itu jika menggunakan concurrency programming bisa digunakan untuk melakukan hal lain

		jadi benefitnya akan bagus sekali untuk jenis seperti ini, dan golang banyak dipake buat backend dengan mendukung concurrency programming dengan konsumsi resource lebih rendah dibanding parallel programming
	*/
}
