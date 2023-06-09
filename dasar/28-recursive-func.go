package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
func main() {
	// recursive func = func yang memanggil func dirinya sendiri
	// contoh kasus mudah diselesaikan dengan recursive adalah factorial
	loop := factorialLoop(5)
	recursive := factorialRecursive(5)
	fmt.Println(loop)
	fmt.Println(recursive)
}
