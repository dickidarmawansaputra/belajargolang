package main

import (
	"fmt"
	"strings"
)

func main() {
	// package strings berisikan func2 untuk memanipulasi data tipe string
	// baca dokumentasi untuk lihat func2 lainnya
	fmt.Println(strings.Contains("Dicki Darmawan Saputra", "Dicki"))
	fmt.Println(strings.ToLower("Dicki Darmawan Saputra"))
	fmt.Println(strings.ToUpper("Dicki Darmawan Saputra"))
}
