package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// disesuaikan kalo masukan text dalam file .txt berarti string
//
//go:embed version.txt
var version string

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
}

//go:embed wallpaper.jpg
var gambar []byte

// untuk embed file binary
func TestEmbedByte(t *testing.T) {
	err := ioutil.WriteFile("wallpaper_new.jpg", gambar, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/1.txt
//go:embed files/2.txt
//go:embed files/3.txt
var files embed.FS

func TestEmbedMultiple(t *testing.T) {
	satu, _ := files.ReadFile("files/1.txt")
	fmt.Println(string(satu))
	dua, _ := files.ReadFile("files/2.txt")
	fmt.Println(string(dua))
	tiga, _ := files.ReadFile("files/3.txt")
	fmt.Println(string(tiga))
}

//go:embed files/*.txt
var path embed.FS // embed.FS = embed filesystem

func TestPathMatcher(t *testing.T) {
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
