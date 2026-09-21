package main

import "fmt"

const NMAX = 999

type song struct {
	judul       string
	penyanyi    string
	durasiMenit int
	durasiDetik int
}

type TabLagu struct {
	totalDurasi int
	arrLagu     [NMAX]song
}

func inputLagu(s *TabLagu, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("Judul lagu: ")
		fmt.Scan(&s.arrLagu[i].judul)
		fmt.Printf("Nama Penyanyi: ")
		fmt.Scan(&s.arrLagu[i].penyanyi)
		fmt.Printf("Durasi lagu (menit detik): ")
		fmt.Scan(&s.arrLagu[i].durasiMenit, &s.arrLagu[i].durasiDetik)
		s.totalDurasi += (s.arrLagu[i].durasiMenit * 60) + s.arrLagu[i].durasiDetik
	}
}

func printLagu(s TabLagu, n int) {
	var i        int
	var jam      int
	var menit    int
	var detikSisa int
	var durasi   string
	var totalStr string

	fmt.Println("+----------------+--------------+----------+")
	fmt.Printf("| %-14s | %-12s | %-8s |\n", "Judul Lagu", "Penyanyi", "Durasi")
	fmt.Println("+----------------+--------------+----------+")

	for i = 0; i < n; i++ {
		durasi = fmt.Sprintf("%02d:%02d",
			s.arrLagu[i].durasiMenit,
			s.arrLagu[i].durasiDetik)
		fmt.Printf("| %-14s | %-12s | %-8s |\n",
			s.arrLagu[i].judul,
			s.arrLagu[i].penyanyi,
			durasi)
	}

	fmt.Println("+----------------+--------------+----------+")

	jam       = s.totalDurasi / 3600
	menit     = (s.totalDurasi % 3600) / 60
	detikSisa = s.totalDurasi % 60
	totalStr  = fmt.Sprintf("%02d:%02d:%02d", jam, menit, detikSisa)
	fmt.Printf("| %-29s | %-8s |\n", "Total Durasi Lagu", totalStr)
	fmt.Println("+----------------+--------------+----------+")
}

func main() {
	var s TabLagu
	var n int

	s.totalDurasi = 0
	fmt.Scan(&n)
	inputLagu(&s, n)
	printLagu(s, n)
}