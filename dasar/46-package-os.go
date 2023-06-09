package main

import (
	"fmt"
	"os"
)

func main() {
	// di go banyak package bawaan
	// packagge os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara idependen
	args := os.Args
	fmt.Println(args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error:", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)

}
