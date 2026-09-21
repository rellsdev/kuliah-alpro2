package main

import "fmt"

const NMAX = 100

type character struct {
	name    string
	health  int
	attacks [NMAX]int
}

func processAttack(penyerang *character, lawan *character, countP *int, dmg int) {
	penyerang.attacks[*countP] = dmg
	*countP = *countP + 1
	lawan.health = lawan.health - dmg
}

func processDefense(bertahan *character, lastDmg int) {
	bertahan.health = bertahan.health + lastDmg
}

func processParry(penyerang *character, lawan *character, countPenyerang int, lastDmgLawan int) {
	penyerang.health = penyerang.health + lastDmgLawan
	lawan.health = lawan.health - penyerang.attacks[countPenyerang-1]
}

func totalDamage(c character, count int) int {
	var i     int
	var total int
	total = 0
	for i = 0; i < count; i++ {
		total = total + c.attacks[i]
	}
	return total
}

func main() {
	var link       character
	var ganon      character
	var linkCount  int
	var ganonCount int
	var namaAksi   string
	var jenis      string
	var damage     int

	link.name  = "Link"
	ganon.name = "Ganon"
	linkCount  = 0
	ganonCount = 0

	fmt.Scan(&link.health, &ganon.health)

	for {
		fmt.Scan(&namaAksi, &jenis)

		if namaAksi == "Link" {
			if jenis == "ATTACK" {
				fmt.Scan(&damage)
				processAttack(&link, &ganon, &linkCount, damage)
			} else if jenis == "DEFENSE" {
				if ganonCount > 0 {
					processDefense(&link, ganon.attacks[ganonCount-1])
				}
			} else if jenis == "PARRY" {
				if ganonCount > 0 {
					processParry(&link, &ganon, linkCount, ganon.attacks[ganonCount-1])
				}
			}
		} else {
			if jenis == "ATTACK" {
				fmt.Scan(&damage)
				processAttack(&ganon, &link, &ganonCount, damage)
			} else if jenis == "DEFENSE" {
				if linkCount > 0 {
					processDefense(&ganon, link.attacks[linkCount-1])
				}
			} else if jenis == "PARRY" {
				if linkCount > 0 {
					processParry(&ganon, &link, ganonCount, link.attacks[linkCount-1])
				}
			}
		}

		// Cek kondisi setelah aksi diproses
		if link.health <= 0 || ganon.health <= 0 {
			break
		}
	}

	if link.health <= 0 {
		fmt.Println("Ganon menang! putri Zelda dalam bahaya.")
	} else {
		fmt.Println("Link menang! putri Zelda berhasil diselamatkan.")
	}

	fmt.Printf("Link: health %d, attacks %d, damage dealt %d\n",
		link.health, linkCount, totalDamage(link, linkCount))
	fmt.Printf("Ganon: health %d, attacks %d, damage dealt %d\n",
		ganon.health, ganonCount, totalDamage(ganon, ganonCount))
}