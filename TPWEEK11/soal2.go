package main

import "fmt"

const NMAX int = 1000

type penduduk struct {
	nama       string
	tahunLahir int
	kota       string
}

type tabPenduduk [NMAX]penduduk

func inputData(A *tabPenduduk, N *int, tahunCari *int) {
	fmt.Scan(N)

	for i := 0; i < *N; i++ {
		fmt.Scan(&A[i].nama, &A[i].tahunLahir, &A[i].kota)
	}

	fmt.Scan(tahunCari)
}

func binarySearch(A tabPenduduk, N int, tahunCari int) int {
	left := 0
	right := N - 1
	idx := -1
	found := false

	for left <= right && !found {
		mid := (left + right) / 2

		if A[mid].tahunLahir == tahunCari {
			found = true
			idx = mid
		} else if A[mid].tahunLahir < tahunCari {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return idx
}

func main() {
	var data tabPenduduk
	var N int
	var tahunCari int

	inputData(&data, &N, &tahunCari)

	idx := binarySearch(data, N, tahunCari)

	if idx != -1 {
		fmt.Println(data[idx].nama)
		fmt.Println(data[idx].tahunLahir)
		fmt.Println(data[idx].kota)
		fmt.Printf("ditemukan di index ke-%d\n", idx)
	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}
