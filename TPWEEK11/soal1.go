package main

import "fmt"

const NMAX int = 1001

type tabMurid [NMAX]string

func inputData(A *tabMurid, N *int, namaCari *string) {
	fmt.Scan(N)

	for i := 0; i < *N; i++ {
		fmt.Scan(&A[i])
	}

	fmt.Scan(namaCari)
}

func binarySearch(A tabMurid, N int, namaCari string) (bool, int) {
	left := 0
	right := N - 1
	posisi := -1
	found := false

	for left <= right && !found {
		mid := (left + right) / 2

		if A[mid] == namaCari {
			found = true
			posisi = mid + 1
		} else if A[mid] < namaCari {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return found, posisi
}

func main() {
	var murid tabMurid
	var N int
	var namaCari string

	inputData(&murid, &N, &namaCari)

	ditemukan, posisi := binarySearch(murid, N, namaCari)

	if ditemukan {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", posisi)
	} else {
		fmt.Println("Murid tidak terdaftar")
	}
}
