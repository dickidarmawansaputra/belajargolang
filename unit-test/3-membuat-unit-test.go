package main

func main() {
	// aturan file test
	/*
		golang memiliki aturan cara membuat file unit test
		untuk membuat unit test, harus menggunakan akhiran _test
		misal membuat file hello_world.go untuk unit testnya jadi hello_world_test.go
	*/

	// aturan nama func unit test
	/*
		nama func untuk unit test diawali nama Test
		misal untuk mengetest func HelloWorld makan func unit testnya TestHelloWorld
		harus memiliki paramater (t *testing.T) dan tidak mengembalikan return value
	*/

	// menjalankan unit test
	/*
		untuk menjalankan unit test dengan perintah: "go test"
		jika ingin lihat detail func test apa saja yang sudah di running dengan perintah: "go test -v"
		jika ingin memilih func unit test mana yg ingin di running dengan perintah: "go test -v -run TestNamaFunc"
		jika ingin menjalankan semua unit test tanpa masuk ke folder testnya/dari root folder dengan perintah: "go test ./..."
	*/
}
