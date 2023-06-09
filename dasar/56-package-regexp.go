package main

import (
	"fmt"
	"regexp"
)

func main() {
	// package regexp = utilitas di go untuk pencarian regular expression
	// expressionnya cek di dokumentasi
	regex := regexp.MustCompile("dic([a-z])i")
	fmt.Println(regex.MatchString("dicki"))    // true
	fmt.Println(regex.MatchString("darmawan")) // false
	fmt.Println(regex.MatchString("saputra"))  // false

	search := regex.FindAllString("dicki docko diki daka", 1)
	fmt.Println(search)
}
