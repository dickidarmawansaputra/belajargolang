package main

import (
	"fmt"
	"strconv"
)

func main() {
	// sebelumnya sudah tau cara konversi tipe data int32 ke int64
	// bagaimana jika konversi yg tipe datanya berbeda misal dari int ke string atau sebaliknya
	// bisa dengan package strconv (string conversion)
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error:", err.Error())
	}

	// basenya 10 = decimal, bisa octal dll tergantung kebutuhan
	number, err := strconv.ParseInt("1000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("Error:", err.Error())
	}

	// basenya 10 = decimal, bisa octal dll tergantung kebutuhan
	value := strconv.FormatInt(1000, 10)
	fmt.Println(value)

	// secara otomatis tanpa tentuin manual basenya misal decimal
	valueString := strconv.Itoa(1000)
	valueInt, err := strconv.Atoi("1000")
	fmt.Println(valueString)
	fmt.Println(valueInt)
}
