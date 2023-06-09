package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}

// validasi required
func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	// fitur reflection = dimana bisa melihat struktur kode pada saat aplikasi sendang berjalan
	// bisa membaca struktur aplikasi saat aplikasi sudah di compile
	// reflection sangat berguna saat ingin membuat library yg general sehingga mudah digunakan
	sample := Sample{"Dicki"}
	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.Field(0))

	// StructTag
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	fmt.Println(IsValid(sample)) // true

	contoh := ContohLagi{"", ""}
	fmt.Println(IsValid(contoh)) // false
}
