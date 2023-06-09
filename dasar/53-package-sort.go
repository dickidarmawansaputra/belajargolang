package main

import (
	"fmt"
	"sort"
)

type Pengguna struct {
	Nama string
	Umur int
}

type PenggunaSlice []Pengguna

func (value PenggunaSlice) Len() int {
	return len(value)
}

func (value PenggunaSlice) Less(i, j int) bool {
	return value[i].Umur < value[j].Umur
}

func (value PenggunaSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	// package sort = berisikan utilitas untuk pengurutan data
	// agar data bisa diurutkan harus implementasikan kontrak sort.Interface
	users := []Pengguna{
		{"a", 12},
		{"b", 20},
		{"c", 18},
		{"d", 15},
		{"e", 10},
	}

	// sort ascending
	sort.Sort(PenggunaSlice(users))
	fmt.Println(users)

	// untuk sort descending bisa reverse seperti ini atau ubah kondisi di func les jadi >
	sort.Sort(sort.Reverse(PenggunaSlice(users)))
	fmt.Println(users)
}
