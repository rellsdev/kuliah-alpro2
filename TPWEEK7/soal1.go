package main

import "fmt"

func validSolution(x int) bool {
	var digits [10]int

	temp := x
	for i := 1; i <= 9; i++ {
		digit := temp % 10
		temp = temp / 10
		if digit == 0 {
			return false
		}
		digits[digit]++
	}

	for i := 1; i <= 9; i++ {
		if digits[i] != 1 {
			return false
		}
	}

	return true
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(validSolution(x))
}
