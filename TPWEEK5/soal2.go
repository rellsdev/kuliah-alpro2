package main

import "fmt"

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func hitungDigitRek(nomor string, posisi int) int {
	if posisi > len(nomor) {
		return 0
	}
	if isDigit(nomor[posisi-1]) {
		return 1 + hitungDigitRek(nomor, posisi+1)
	}
	return hitungDigitRek(nomor, posisi+1)
}

func awalanValid(nomor string) bool {
	if len(nomor) >= 3 && nomor[0] == '+' && nomor[1] == '6' && nomor[2] == '2' {
		return true
	}
	if len(nomor) >= 2 && nomor[0] == '0' && nomor[1] == '8' {
		return true
	}
	return false
}

func validasiNomor(nomor string) bool {
	if !awalanValid(nomor) {
		return false
	}
	var startPos int
	if nomor[0] == '+' {
		startPos = 4 // setelah "+62"
	} else {
		startPos = 3 // setelah "08"
	}
	var jumlahDigit int = hitungDigitRek(nomor, startPos)
	return jumlahDigit >= 9 && jumlahDigit <= 12
}

func main() {
	var nomor string
	fmt.Scan(&nomor)

	if validasiNomor(nomor) {
		fmt.Println("VALID")
	} else {
		fmt.Println("TIDAK VALID")
	}
}