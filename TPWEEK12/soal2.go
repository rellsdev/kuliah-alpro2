package main

import "fmt"

const NMAX int = 9999

type arrString [NMAX]string

func bacaData(A *arrString, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A arrString, n int) {
	for i := 0; i < n; i++ {
		if i < n-1 {
			fmt.Printf("%s, ", A[i])
		} else {
			fmt.Printf("%s.", A[i])
		}
	}
	fmt.Println()
}

func selectionSortAscend(A *arrString, n int) {
	var pass, idx, i int
	var temp string
	for pass = 1; pass <= n-1; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if A[i] < A[idx] {
				idx = i
			}
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
	}
}

func selectionSortDescend(A *arrString, n int) {
	var pass, idx, i int
	var temp string
	for pass = 1; pass <= n-1; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if A[i] > A[idx] {
				idx = i
			}
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
	}
}

func main() {
	var A arrString
	var n int

	fmt.Scan(&n)
	bacaData(&A, n)

	selectionSortAscend(&A, n)
	fmt.Print("Data setelah diurutkan secara Ascending:\n")
	cetakData(A, n)

	selectionSortDescend(&A, n)
	fmt.Print("Data setelah diurutkan secara Descending:\n")
	cetakData(A, n)
}
