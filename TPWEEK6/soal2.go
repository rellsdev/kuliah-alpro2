package main

import "fmt"

func hariKerja(hariMulai, jumlahHari int) int {
	if jumlahHari == 0 {
		return 0
	}
	if hariMulai >= 1 && hariMulai <= 5 {
		return 1 + hariKerja((hariMulai%7)+1, jumlahHari-1)
	}
	return hariKerja((hariMulai%7)+1, jumlahHari-1)
}

func main() {
	var hariMulai, jumlahHari int
	fmt.Scan(&hariMulai, &jumlahHari)
	fmt.Println(hariKerja(hariMulai, jumlahHari))
}
