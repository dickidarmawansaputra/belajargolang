package main

import (
	"container/list"
	"fmt"
)

func main() {
	// package container/list = implementasi struktur data double linked list di go
	data := list.New()

	data.PushBack("Dicki")
	data.PushBack("Dicki Darmawan")
	data.PushBack("Dicki Darmawan Saputra")
	data.PushFront("Roy")

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Front().Next().Next().Value)
	// fmt.Println(data.Back().Prev().Value)
	// fmt.Println(data.Back().Value)

	// dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
