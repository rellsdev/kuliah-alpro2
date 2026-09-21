package main

import "fmt"

const NMAX = 999

type menu struct {
	nama   string
	harga  int
	stok   int
	status string
}

type arrMenu [NMAX]menu

func inputDataMenu(m *arrMenu, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("Nama Menu: ")
		fmt.Scan(&m[i].nama)
		fmt.Printf("Harga Menu: Rp.")
		fmt.Scan(&m[i].harga)
		fmt.Printf("Jumlah Stok: ")
		fmt.Scan(&m[i].stok)
		if m[i].stok > 0 {
			m[i].status = "Tersedia"
		} else {
			m[i].status = "Kosong"
		}
	}
}

func printDataMenu(m arrMenu, n int) {
	var i int
	fmt.Printf("%-15s %-12s %-11s %-s\n",
		"Nama Menu", "Harga (Rp)", "Stok (pcs)", "Status Kesediaan")
	for i = 0; i < n; i++ {
		fmt.Printf("%-15s Rp.%-9d %-11d %-s\n",
			m[i].nama,
			m[i].harga,
			m[i].stok,
			m[i].status)
	}
}

func main() {
	var m arrMenu
	var n int

	fmt.Scan(&n)
	inputDataMenu(&m, n)
	printDataMenu(m, n)
}