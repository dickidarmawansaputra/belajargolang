package main

import (
	"flag"
	"fmt"
)

func main() {
	// package flag berisikan fungsionalitas untuk memparsing command line argument
	// memudahkan jika ingin bikin aplikasi berbasis terminal
	host := flag.String("hostname", "localhost", "put your host")
	user := flag.String("username", "root", "put your host")
	pass := flag.String("password", "root", "put your host")

	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("User:", *user)
	fmt.Println("Pass:", *pass)
}
