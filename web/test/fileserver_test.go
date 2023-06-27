package test

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestFileServer(t *testing.T) {
	// jika di build production (dicompile) di server file di folder resource harus diupload juga (jadi 2 kali kerja, ribet)
	// solusinya pake golang embed agar ketika di compile tidak perlu lagi upload filenya di server
	// liat di test file server golang embed
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	// Hal ini dikarenakan FileServer akan membaca url, lalu mencari file berdasarkan url nya, jadi jika
	// kita membuat /static/index.js, maka FileServer akan mencari ke file /resources/static/index.js
	// solusi 404 not found dengan http.StripPrefix
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestFileServerGoEmbed(t *testing.T) {
	/*
		Kenapa ini terjadi? Hal ini karena di Go-Lang embed, nama folder ikut menjadi nama resource nya,
		misal resources/index.js, jadi untuk mengaksesnya kita perlu gunakan URL
		/static/resources/index.js
	*/
	// solusinya agar tidak 404 Not Found dengan fs.Sub() func
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
