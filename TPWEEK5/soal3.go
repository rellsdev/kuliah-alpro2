package main

import "fmt"

func hitungPoin(belanja int) int {
	if belanja < 50000 {
		return 0
	} else if belanja < 100000 {
		return 5
	} else if belanja < 200000 {
		return 10
	} else {
		return 20
	}
}

func prosesTransaksi(nomorTransaksi int, totalPoin int) {
	var belanja int
	fmt.Scan(&belanja)

	if belanja == 0 {
		fmt.Println("Total Poin:", totalPoin)
	} else {
		var poin int = hitungPoin(belanja)
		fmt.Printf("Transaksi %d: Belanja Rp %d, Poin: %d\n",
			nomorTransaksi, belanja, poin)
		prosesTransaksi(nomorTransaksi+1, totalPoin+poin)
	}
}

func main() {
	prosesTransaksi(1, 0)
}
