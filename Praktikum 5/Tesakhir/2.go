package main

import "fmt"

func main() {
	var tendangan, gol int
	var i int

	gol = 0
	i = 1
	fmt.Print("Tendangan ke-", i, " : ")
	fmt.Scan(&tendangan)
	for tendangan > 0 {
		i++
		fmt.Print("Tendangan ke-", i, " : ")
		fmt.Scan(&tendangan)
		if tendangan % 4 == 0 {
			if tendangan != 0 {
				gol++
			}
		}	
	}
	i--
	if gol == (i / 2) {
		fmt.Println("Draw")
	} else if gol > (i / 2) {
		fmt.Println("Penyerang menang")	
	} else {
		fmt.Println("Kiper menang")
	}
	fmt.Printf("Skor gol: %d dari %d tendangan", gol, i)
}