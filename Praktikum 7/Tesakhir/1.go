package main

import "fmt"

func tanggunganHari(jumHari int, tujuan string) int {
	if tujuan == "domestik" || tujuan == "Domestik"{
		if jumHari > 3 {
			jumHari = 3
			return jumHari
		}else{
			return jumHari
		}
	}else{
		if jumHari > 8 {
			jumHari = 8
			return jumHari
		}else{
			return jumHari
		}
	}
}

func biayaPerhari(jumMahasiswa int, jumHari int, tujuan string) int {
	return (350 + 2500 + 3000) * jumMahasiswa * tanggunganHari(jumHari, tujuan)
}

func perhitunganBiaya(jumMahasiswa int, jumHari int, tujuan string, totalBiaya *int){
	if tujuan == "domestik" || tujuan == "Domestik"{
		*totalBiaya = biayaPerhari(jumMahasiswa, jumHari, tujuan)
	}else{
		*totalBiaya = 15 * biayaPerhari(jumMahasiswa, jumHari, tujuan) / 10
	}
}

func main(){
	var jumMahasiswa, jumHari int
	var tujuan string
	var totalBiaya int

	totalBiaya = 0
	fmt.Scan(&jumMahasiswa, &jumHari, &tujuan)
	perhitunganBiaya(jumMahasiswa, jumHari, tujuan, &totalBiaya)
	fmt.Print(totalBiaya)
}