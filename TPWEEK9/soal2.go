package main

import "fmt"

const NMAX int = 100

type item struct {
	nama  string
	harga int
}

type tabItem [NMAX]item

func isiData(t *tabItem, n int) {
	var i int

	i = 0
	for i < n {
		fmt.Print("Nama item: ")
		fmt.Scan(&t[i].nama)

		fmt.Print("Harga item: ")
		fmt.Scan(&t[i].harga)

		i = i + 1
	}
}

func tampilkanData(t tabItem, n int) {
	var i int

	fmt.Println("Nama Item\t\tHarga")

	i = 0
	for i < n {
		fmt.Printf("%-20s\t%d\n", t[i].nama, t[i].harga)
		i = i + 1
	}
}

func cariTermahal(t tabItem, n int) int {
	var idxMax int
	var i int

	idxMax = 0
	i = 1

	for i < n {
		if t[idxMax].harga < t[i].harga {
			idxMax = i
		}
		i = i + 1
	}

	return idxMax
}

func cariTermurah(t tabItem, n int) int {
	var idxMin int
	var i int

	idxMin = 0
	i = 1

	for i < n {
		if t[idxMin].harga > t[i].harga {
			idxMin = i
		}
		i = i + 1
	}

	return idxMin
}

func main() {
	var data tabItem
	var n int
	var idxMahal int
	var idxMurah int

	fmt.Scan(&n)

	isiData(&data, n)

	fmt.Println()
	tampilkanData(data, n)

	idxMahal = cariTermahal(data, n)
	idxMurah = cariTermurah(data, n)

	fmt.Println()
	fmt.Println("Item termahal", data[idxMahal].nama)
	fmt.Println("Item termurah", data[idxMurah].nama)
}
