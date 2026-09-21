package main

import "fmt"

const NMAX = 100

type Peserta struct {
	id     string
	nama   string
	nilai  int
	durasi int
}

type TabPeserta [NMAX]Peserta

func isiArray(T *TabPeserta, N int) {
	for i := 0; i < N; i++ {
		fmt.Scan(&T[i].id, &T[i].nama, &T[i].nilai, &T[i].durasi)
	}
}

func insertionSort(T *TabPeserta, N int) {
	for pass := 1; pass < N; pass++ {
		temp := T[pass]
		i := pass

		for i > 0 && (temp.nilai > T[i-1].nilai ||
			(temp.nilai == T[i-1].nilai && temp.durasi < T[i-1].durasi)) {

			T[i] = T[i-1]
			i--
		}

		T[i] = temp
	}
}

func tampilArray(T TabPeserta, N int) {
	for i := 0; i < N; i++ {
		fmt.Println(T[i].id, T[i].nama, T[i].nilai, T[i].durasi)
	}
}

func tampilTerbaik(T TabPeserta) {
	fmt.Println(T[0].id, T[0].nama, T[0].nilai, T[0].durasi)
}

func hitungDiAtasRata(T TabPeserta, N int) {
	total := 0
	for i := 0; i < N; i++ {
		total += T[i].nilai
	}

	rata := float64(total) / float64(N)

	count := 0
	for i := 0; i < N; i++ {
		if float64(T[i].nilai) > rata {
			count++
		}
	}

	fmt.Println(count)
}

func main() {
	var T TabPeserta
	var N int

	fmt.Scan(&N)

	isiArray(&T, N)

	insertionSort(&T, N)

	fmt.Println("Data setelah diurutkan:")
	tampilArray(T, N)

	fmt.Println("\nPeserta terbaik:")
	tampilTerbaik(T)

	fmt.Print("\nJumlah peserta di atas rata-rata: ")
	hitungDiAtasRata(T, N)
}
