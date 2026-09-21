package main

import "fmt"

const NMAX int = 100

type applicant struct {
	id           string
	name         string
	testGrade    float64
	testDuration int
}

type applicants struct {
	tab [NMAX]applicant
	n   int
}

func inputData(app *applicants) {
	var id string
	for {
		fmt.Scan(&id)
		if id == "END" {
			break
		}
		var name string
		var grade float64
		var duration int
		fmt.Scan(&name, &grade, &duration)
		app.tab[app.n] = applicant{id, name, grade, duration}
		app.n++
	}
}

func printData(app applicants) {
	var maxLen int
	for i := 0; i < app.n; i++ {
		if len(app.tab[i].name) > maxLen {
			maxLen = len(app.tab[i].name)
		}
	}

	for i := 0; i < app.n; i++ {
		fmt.Printf("%-6s %-*s %.1f %d\n",
			app.tab[i].id,
			maxLen,
			app.tab[i].name,
			app.tab[i].testGrade,
			app.tab[i].testDuration)
	}
}

func shouldSwap(app applicants, idx, i int) bool {
	a := app.tab[idx]
	b := app.tab[i]
	if b.testGrade > a.testGrade {
		return true
	}
	if b.testGrade == a.testGrade && b.testDuration < a.testDuration {
		return true
	}
	return false
}

func sortApplicants(app *applicants) {
	var pass, idx, i int
	var temp applicant
	for pass = 1; pass <= app.n-1; pass++ {
		idx = pass - 1
		for i = pass; i < app.n; i++ {
			if shouldSwap(*app, idx, i) {
				idx = i
			}
		}
		temp = app.tab[pass-1]
		app.tab[pass-1] = app.tab[idx]
		app.tab[idx] = temp
	}
}

func calcAverage(app applicants) (float64, float64) {
	var totalGrade float64
	var totalDuration int
	for i := 0; i < app.n; i++ {
		totalGrade += app.tab[i].testGrade
		totalDuration += app.tab[i].testDuration
	}
	return totalGrade / float64(app.n),
		float64(totalDuration) / float64(app.n)
}

func main() {
	var app applicants
	app.n = 0

	inputData(&app)

	fmt.Println("Data awal:")
	printData(app)

	sortApplicants(&app)

	fmt.Println("\nData setelah sortir:")
	printData(app)

	avgGrade, avgDuration := calcAverage(app)
	fmt.Printf("\nRata-rata nilai tes: %.2f\n", avgGrade)
	fmt.Printf("Rata-rata durasi: %.2f\n", avgDuration)
}
