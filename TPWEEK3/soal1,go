package main

import "fmt"

func reverseAngka(x int) int {
	hasil := 0
	for x > 0 {
		hasil = hasil*10 + (x % 10)
		x /= 10
	}
	return hasil
}

func ambilGabungkanTigaDigit(x int) int {
	reversed := reverseAngka(x)       
	tigaDigit := reversed % 1000      
	sisa := reversed / 1000           
	tigaDigitReverse := reverseAngka(tigaDigit)
	return sisa*1000 + tigaDigitReverse
}


func cekGanjilGenap(x int) {
	ganjil, genap := 0, 0
	for x > 0 {
		digit := x % 10
		if digit%2 == 0 {
			genap++
		} else {
			ganjil++
		}
		x /= 10
	}
	if ganjil > genap {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}

func prosesPola(n int) {
	hasil := ambilGabungkanTigaDigit(n)
	fmt.Println(hasil)
	cekGanjilGenap(hasil)
}

func main() {
	var n int
	fmt.Scan(&n)
	prosesPola(n)
}