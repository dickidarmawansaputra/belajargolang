package main

import "fmt"

// jika lebih dari 1 argument varargs harus dibelakang
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func main() {
	// parameter di posisi terakhir memiliki kemampuan jadi sebuah varargs
	// varargs bisa nerima lebih dari satu input anggap seperti array
	total := sumAll(10, 1, 2, 3, 4)
	fmt.Println(total)

	// masukan data slice ke variadic func
	slice := []int{1, 2, 3, 4, 5}
	totalSum := sumAll(slice...)
	fmt.Println(totalSum)
}
