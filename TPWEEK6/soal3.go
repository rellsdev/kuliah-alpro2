package main

import "fmt"

func isPeak(n int) bool {
	if n < 100 {
		return false
	}
	var kanan int = n % 10
	var tengah int = (n / 10) % 10
	var kiri int = (n / 100) % 10
	if tengah > kiri && tengah > kanan {
		return true
	}
	return isPeak(n / 10)
}

func main() {
	var n int
	fmt.Scan(&n)
	if isPeak(n) {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}
