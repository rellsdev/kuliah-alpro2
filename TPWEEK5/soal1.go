package main

import "fmt"

func countDigit(n int) int {
	if n < 10 {
		return 1
	}
	return 1 + countDigit(n/10)
}

func getNumber(n, i int) int {
	var total int = countDigit(n)
	if i == total {
		return n % 10
	}
	return getNumber(n/10, i)
}

func checkFirstLast(n int) bool {
	var total int = countDigit(n)
	var first int = getNumber(n, 1)
	var last int = getNumber(n, total)
	return first == last
}

func sumDigit(n int) int {
	if n < 10 {
		return n
	}
	return (n % 10) + sumDigit(n/10)
}

func isPrima(n, i int) bool {
	if n < 2 {
		return false
	}
	if i*i > n {
		return true
	}
	if n%i == 0 {
		return false
	}
	return isPrima(n, i+1)
}

func main() {
	var N int
	fmt.Scan(&N)

	var total int = countDigit(N)

	if N > 100 && total%2 == 1 {
		var tengah int = (total + 1) / 2
		var digitTengah int = getNumber(N, tengah)
		var jumlah int = sumDigit(N)

		var tengahPrima bool = isPrima(digitTengah, 2)
		var samaDep bool = checkFirstLast(N)
		var kelipatanTiga bool = jumlah%3 == 0

		if tengahPrima && samaDep && kelipatanTiga {
			fmt.Println("Pesawat merupakan pesawat spesial.")
		} else if !samaDep {
			fmt.Println("Pesawat merupakan pesawat berisiko.")
		} else if !tengahPrima {
			fmt.Println("Pesawat merupakan pesawat biasa.")
		} else {
			fmt.Println("Pesawat tidak layak terbang.")
		}
	} else {
		fmt.Println("Bilangan tidak valid.")
	}
}