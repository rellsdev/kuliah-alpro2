package main

import "fmt"

const NMAX = 100

type arrInt [NMAX]int

func digitPuluhan(x int) int {
	return (x / 10) % 10
}

func tampilArray(A arrInt, N int) {
	for i := 0; i < N; i++ {
		fmt.Println(A[i])
	}
}

func insertionSortModifikasi(A *arrInt, N int) {
	for pass := 1; pass < N; pass++ {
		temp := A[pass]
		i := pass

		for i > 0 && (digitPuluhan(temp) < digitPuluhan(A[i-1]) ||
			(digitPuluhan(temp) == digitPuluhan(A[i-1]) && temp < A[i-1])) {

			A[i] = A[i-1]
			i--
		}

		A[i] = temp
	}
}

func main() {
	var A arrInt
	var N int

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	fmt.Println("Data sebelum sorting:")
	tampilArray(A, N)

	insertionSortModifikasi(&A, N)
	fmt.Println()

	fmt.Println("Data setelah sorting:")
	tampilArray(A, N)
}
