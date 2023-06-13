package main

func main() {
	// saat bikin aplikasi biasa menggunakan library atau dependency dari project lain
	// sebelum ada go modules, management dependency di go sangat sulit
	//  biasanya copy semua source code library atau dependency ke project kita & project jadi bengkak untuk sizenya
	// go module mirip composer jika di php

	// untuk buat module dengan perintah "go mod ini nama module"
	// secara otomatis membuat file go.mod berisikan info nama module dan versi go yg digunakan
	// go terintegrasi baik dengan git
	// untuk rilis module perlu membuat tag di git

	// membuat module
	/*
		1. create folder project
		2. create repo di github
		3. jalankan perintah "go mod init nama module. example: go mod init github.com/dickidarmawansaputra/go-module-sayhello"
		4. git add & commit & push ke repo
		5. git tag dengan perintah "git tag v1.0.0" & "git push origin v1.0.0"
	*/

	// menambahkan dependency
	/*
		1. create folder project
		2. create repo di github
		3. jalankan perintah "go get nama module. example: go get github.com/dickidarmawansaputra/go-module-sayhello"
		4. tinggal pangil nama func yang ada di module
		5. push code ke repo
	*/

	// upgrade module
	/*
		1. cukup membuat tag baru & releas tag
	*/

	// upgrade dependency
	/*
		1. untuk upgrade ubah isi go.mod ubah tag modulenya ke versi baru
		2. jalankan perintah "go get"
	*/

	// major upgrade di module
	// major upgrade terjadi dikarenakan ada perubahan pada isi kode program sehingga tidak backward compatible
	// sebaiknya sebisa mungkin dihindari
	// namun jika tidak bisa dihindari, strategy terbaik dengan merubah nama module
	/*
		1. merubah nama module dengan tambahan versi modulenya di file go.mod jika ada major changes di module
		2. example: sebelumnya nama module "github.com/dickidarmawansaputra/go-module-sayhello" jadi "github.com/dickidarmawansaputra/go-module-sayhello/v2" di file go.mod
		3. git commit & push kemudian buat git tag versi baru
	*/
	// cek di go fiber salah satu framework popular golang
}
