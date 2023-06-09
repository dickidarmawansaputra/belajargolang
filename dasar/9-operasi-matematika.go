package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var x = 10
	var y = 15
	var z = x + y
	fmt.Println(z)

	// augmented assignment
	var i = 10

	// artinya i = i + 10
	i += 10
	fmt.Println(i)

	// unary operator
	i++
	fmt.Println(i)

	var negative = -100
	fmt.Println(negative)
}
