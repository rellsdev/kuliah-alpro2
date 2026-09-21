package main

import "fmt"

const NMAX int = 100

type Student struct {
	nim   string
	name  string
	grade float64
}

type Students [NMAX]Student

func cariNIM(A Students, n int, nim string) int {
	var i int

	i = 0
	for i < n {
		if A[i].nim == nim {
			return i
		}
		i++
	}

	return -1
}

func inputDataUnik(A *Students, n *int) {
	var nim, name string
	var grade float64

	*n = 0

	fmt.Scan(&nim)
	for nim != "STOP" {
		fmt.Scan(&name, &grade)

		if cariNIM(*A, *n, nim) == -1 {
			A[*n].nim = nim
			A[*n].name = name
			A[*n].grade = grade
			*n++
		}

		fmt.Scan(&nim)
	}
}

func tampilData(A Students, n int) {
	var i int

	i = 0
	for i < n {
		fmt.Println(A[i].nim, A[i].name, A[i].grade)
		i++
	}
}

func main() {
	var A Students
	var n int

	inputDataUnik(&A, &n)
	tampilData(A, n)
}
