package main

import "fmt"

const NMAX int = 999

type Rapot struct {
	matkul string
	nilai  int
}

type TabRapot [NMAX]Rapot

func isiRapot(A *TabRapot, N int) {
	var i int

	i = 0
	for i < N {
		fmt.Scan(&A[i].matkul, &A[i].nilai)
		i++
	}
}

func seqSearch(A TabRapot, N int, X int) bool {
	var i int
	var found bool

	found = false
	i = 0

	for i < N {
		if A[i].nilai == X {
			fmt.Println(A[i].matkul, A[i].nilai)
			found = true
		}
		i++
	}

	return found
}

func main() {
	var A TabRapot
	var N, X int
	var ketemu bool

	fmt.Scan(&N)

	isiRapot(&A, N)

	fmt.Scan(&X)

	ketemu = seqSearch(A, N, X)

	if ketemu {
		fmt.Println("Data ditemukan!")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
