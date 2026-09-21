package main

import "fmt"

func jumlahDigit(n int) int {
	if n < 10 {
		return 1
	}
	return 1 + jumlahDigit(n/10)
}

func hapusGanjil(n, curr, del int) int {
	if n == 0 {
		return 0
	}
	var digit int = n % 10
	var sisa int = n / 10
	if curr == del {
		return hapusGanjil(sisa, curr+1, del)
	}
	return digit + hapusGanjil(sisa, curr+1, del)*10
}

func hapusGenap(n, curr, del1, del2 int) int {
	if n == 0 {
		return 0
	}
	var digit int = n % 10
	var sisa int = n / 10
	if curr == del1 || curr == del2 {
		return hapusGenap(sisa, curr+1, del1, del2)
	}
	return digit + hapusGenap(sisa, curr+1, del1, del2)*10
}

func main() {
	var N int
	fmt.Scan(&N)

	var panjang int = jumlahDigit(N)

	if panjang <= 2 {
		fmt.Println("INPUT ERROR")
	} else if panjang%2 == 1 {
		var tengah int = (panjang + 1) / 2
		var hasil int = hapusGanjil(N, 1, tengah)
		fmt.Println(hasil)
	} else {
		var del1 int = panjang / 2
		var del2 int = (panjang / 2) + 1
		var hasil int = hapusGenap(N, 1, del1, del2)
		fmt.Println(hasil)
	}
}
