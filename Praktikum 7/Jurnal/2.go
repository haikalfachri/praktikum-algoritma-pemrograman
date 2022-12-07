package main

import "fmt"

func BacaData(usaha *int, nDoa *int, doaOrtu *bool, nilai *string){
	fmt.Print("Banyaknya usaha? ")
	fmt.Scan(&*usaha)
	fmt.Print("Banyaknya doa? ")
	fmt.Scan(&*nDoa)
	fmt.Print("Doa orang tua? ")
	fmt.Scan(&*doaOrtu)
	fmt.Print("Nilai Algoritma? ")
	fmt.Scan(&*nilai)
}

func TabungUsahaDoa(usaha int, doa int, total *int){
	*total = usaha + doa
}

func TabungDoaOrtu(doa bool, total int) int {
	if doa {
		return total * 2
	}
	return total
}

func HasilNilaiAlpro(nilai string, total *int){
	if nilai == "A"{
		*total -= 150
	}else if nilai == "B" {
		*total -= 130
	}else if nilai == "C"{
		*total -= 100
	}
}

func HasilAkhir(poin int) string {
	if poin >= 130 {
		return "Lulus langsung dapat kerja gaji 2 digit"
	}else if poin >= 50 {
		return "Lulus langsung dapat kerja"
	}
	return "Jangan lelah berdoa dan berusaha, tidak ada yang sia-sia dari usaha dan doa"
}

func main(){
	var total, usaha, nDoa int
	var doaOrtu bool
	var nilai string

	BacaData(&usaha, &nDoa, &doaOrtu, &nilai)
	TabungUsahaDoa(usaha, nDoa, &total)
	total = TabungDoaOrtu(doaOrtu, total)
	HasilNilaiAlpro(nilai, &total)
	fmt.Print(HasilAkhir(total))
}