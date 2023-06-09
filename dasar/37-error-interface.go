package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}
func main() {
	// di go sudah ada kontrak default untuk error
	// di go sudah menyediakan library untuk membuat helper secara mudah terdapat di package errors
	result, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("Error", err.Error())
	}
}
