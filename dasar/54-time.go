package main

import (
	"fmt"
	"time"
)

func main() {
	// package time = berisikan fungsionalitas untuk management waktu
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2023, 01, 01, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	// di go agak berbeda formatnya liat di packagenya. contoh seperti ini
	layout := "2006-01-02" // time.RFC3339
	value := now.Format("2006-01-02")

	parse, _ := time.Parse(layout, value)
	fmt.Println(value)
	fmt.Println(parse)
}
