package main

import "fmt"

func validateNumber(x int) {
	var digits [20]int
	n := 0
	temp := x
	for temp > 0 {
		n++
		temp /= 10
	}
	temp = x
	for i := n; i >= 1; i-- {
		digits[i] = temp % 10
		temp /= 10
	}

	pos := [10]int{}
	for i := 0; i < 10; i++ {
		pos[i] = -1
	}
	for i := 1; i <= n; i++ {
		d := digits[i]
		if pos[d] == -1 {
			pos[d] = i
		}
	}

	valid := true

	if pos[3] != -1 && pos[7] != -1 {
		if pos[3] > pos[7] {
			valid = false
		}
	}

	if pos[5] != -1 && pos[0] != -1 {
		if pos[5] > pos[0] {
			valid = false
		}
	}

	if pos[2] != -1 && pos[6] != -1 {
		if pos[2] > pos[6] {
			valid = false
		}
	}

	if valid {
		fmt.Println("BILANGAN VALID")
	} else {
		fmt.Println("BILANGAN TIDAK VALID")
	}
}

func main() {
	var x int
	fmt.Scan(&x)
	validateNumber(x)
}
