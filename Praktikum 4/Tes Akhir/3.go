package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var seed int64
	var hasil string
	var anda, dadang, nilai_dadu int

	fmt.Print("Angka rahasia: ")
	fmt.Scanln(&seed)
	rand.Seed(seed)
	nilai_dadu = rand.Intn(6) + 1
	dadang = rand.Intn(6) + 1
	fmt.Print("Anda: ")
	fmt.Scanln(&anda)
	fmt.Println("Dadang: ", dadang)
	if anda == nilai_dadu {
		hasil = "Pemenang adalah anda"
	} else if dadang == nilai_dadu {
		hasil = "Pemenang adalah Dadang"
	} else if anda == nilai_dadu && dadang == nilai_dadu {
		hasil = "Seri"
	} else {
		hasil = "Tidak ada pemenang"
	}
	fmt.Printf("Nilai dadu %d, %s\n", nilai_dadu, hasil)
}
