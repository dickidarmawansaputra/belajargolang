package main

import "fmt"

func main() {
	// perulangan di golang cuma ada satu, yaitu for loops
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	// perulangan secara manual
	slice := []string{"Dicki", "Darmawan", "Saputra"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range
	for index, value := range slice {
		fmt.Println("index", index, "=", value)
	}

	// jika ingin println valuenya saja, karna di golang jika ada variable yang tidak digunakan akan error
	// maka gunakan kata kunci _ jika tidak ingin menggunakannnya
	for _, value := range slice {
		fmt.Println(value)
	}

	// for range dengan data map
	person := make(map[string]string)
	person["name"] = "Dicki"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println("Key", key, "=", value)
	}
}
