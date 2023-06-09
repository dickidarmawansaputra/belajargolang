package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// func tanpa nama
	// func di dalam variable
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)

	// langsung di dalam func / func dalam parameter
	registerUser("dicki", func(name string) bool {
		return name == "root"
	})

}
