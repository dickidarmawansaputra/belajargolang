package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed wallpaper.jpg
var gambar []byte

//go:embed files/*.txt
var path embed.FS // embed.FS = embed filesystem

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("wallpaper_new.jpg", gambar, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		// cek direktory atau bukan
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
