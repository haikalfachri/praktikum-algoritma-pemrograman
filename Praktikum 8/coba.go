package main

import "fmt"

func hitungSkor(soal, skor *int){
	var waktu int
	
	waktu = 0
	for i:=1;i<=8;i++{
		fmt.Scan(&*skor)
		if *skor != 301{
			*soal++
			waktu += *skor
			*skor = waktu
		}else{
			*soal += 0
			waktu += 0
			*skor = waktu
		}
	}
	
}

func main(){
	var nama, pemenang string
	var skor, soal, tempSkor, tempSoal int
	var skorMenang, soalMenang, waktu int

	tempSkor = 0
	tempSoal = 0
	soal = 0
	fmt.Scan(&nama)
	for nama != "Selesai"{
		hitungSkor(&soal, &skor)
		fmt.Println(soal, skor)
		if waktu < 301{
			if soal > tempSoal{
				pemenang = nama
				skorMenang = skor
				soalMenang = soal
			}else if soal == tempSoal{
				if waktu < tempSkor{
					pemenang = nama
					skorMenang = skor
					soalMenang = soal
				}
			}
		}		
		tempSkor = waktu
		tempSoal = soal
		soal = 0
		fmt.Scan(&nama)
	}
	fmt.Printf("%s %d %d", pemenang, soalMenang, skorMenang)
}