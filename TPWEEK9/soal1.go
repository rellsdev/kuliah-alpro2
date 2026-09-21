package main

import "fmt"

const NMAX int = 20

type game struct {
	nama     string
	populasi float64
	nilai    int
}

type tabGame [NMAX]game

func isiData(t *tabGame, n int) {
	var i int

	i = 0
	for i < n {
		fmt.Scan(&t[i].nama)
		i = i + 1
	}
}

func hitungData(t *tabGame, n int, k *int) {
	var i int
	var pendudukKota int

	i = 0
	for i < n {
		pendudukKota = int(float64(*k) * 0.10)

		t[i].populasi = float64(pendudukKota)

		*k = *k - pendudukKota

		t[i].nilai = pendudukKota % 1000

		i = i + 1
	}
}

func tampilkan(t tabGame, n int) {
	var i int
	var tingkat float64

	i = 0
	for i < n {
		tingkat = float64(t[i].nilai) * 0.025

		if tingkat > 5 {
			fmt.Printf("%s dengan tingkat kejahatan: %.2f%%\n", t[i].nama, tingkat)
		}

		i = i + 1
	}
}

func main() {
	var data tabGame
	var n int
	var k int

	fmt.Scan(&n, &k)

	isiData(&data, n)
	hitungData(&data, n, &k)
	tampilkan(data, n)
}
