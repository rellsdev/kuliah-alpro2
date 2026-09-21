package main

import (
	"fmt"
	"math"
)

func sumAllDigit(bilangan int) int {
	total := 0
	for bilangan > 0 {
		total += bilangan % 10
		bilangan /= 10
	}
	return total
}

func sumUntilBeforeLimit(bil1, bil2, limit int) int {
	for bil1+bil2 < limit {
		bil1 += bil2
	}
	return bil1
}

func isPrima(bilangan int) bool {
	if bilangan < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(bilangan))); i++ {
		if bilangan%i == 0 {
			return false
		}
	}
	return true
}

func digitPertama(n int) int {
	for n >= 10 {
		n /= 10
	}
	return n
}

func bentukBilangan(n int) int {
	if n < 100 {
		lastDigit := n % 10
		if lastDigit == 0 {
			return sumUntilBeforeLimit(n, digitPertama(n), 1000)
		}
		return sumUntilBeforeLimit(n, lastDigit, 1000)
	} else if n < 1000 {
		return sumUntilBeforeLimit(n, sumAllDigit(n), 10000)
	}
	return n
}

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		bilBaru := bentukBilangan(n)
		fmt.Printf("%d --> %d\n", n, bilBaru)
		if isPrima(bilBaru) {
			fmt.Println("PRIMA")
		} else {
			fmt.Println("BUKAN PRIMA")
		}
	}
}