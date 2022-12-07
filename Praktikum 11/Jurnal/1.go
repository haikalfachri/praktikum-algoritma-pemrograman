package main

import (
	"fmt"
)

func main(){
	const NMAX int : 1000

	type lagu struct{
		judul string
		penyanyi string
		durasi waktu
	}
	type waktu struct{
		menit int
		detik int
	}
	type playlist [NMAX]lagu

	var n int
	var list playlist
	
	buatPlaylist(list, n)
	cetakPlaylist(list, n)

}

func cariLagu(T playlist, n int) lagu{
	var input lagu

	fmt.Scan(&input.judul, &input.penyanyi)
	for i:=0;i<n;i++{
		if T[i] == input.judul && T[i] == input.penyanyi{
			return T[i]
		}
	}
	return nil
}

func cariLaguTerlama(T playlist, n int) lagu{
	var tempMenit int
	var tempDetik int
	var laguTerlama lagu

	for i:=0;i<n;i++{
		if T[i].durasi.menit >= TempMenit{
			if T[i].menit > tempMenit{
				laguTerlama = T[i]
			}else if T[i].menit == tempMenit{
				if T[i].detik > tempDetik{
					laguTerlama = T[i]
				}
			}
		} 
		tempMenit = T[i].durasi.menit
		tempDetik = T[i].durasi.detik
	}
	return laguTerlama
}

func buatPlaylist(T *playlist, n *int){
	var input lagu
	var tempJudul string

	fmt.Scan(&input.judul, &input.penyanyi)
	for input.judul != "#" && input.penyanyi != "#"{
		if tempJudul != input.judul{
			fmt.Scan(&input.durasi.menit, &input.durasi.detik)
			T[*n] = input
			*n++
			tempJudul = input.judul
			fmt.Scan(&input.judul, &input.penyanyi)
		}
	}	
}

func cetakPlaylist(T *playlist, n *int){
	var laguTerlama lagu

	laguTerlama = cariLaguTerlama(T, n)
	for i:=0;i<*n;i++{
		if laguTerlama == T[i]{
			fmt.Printf("*%v %v menit %v detik\n", laguTerlama, T[i].durasi.menit, T[i].durasi.detik)
		}else{
			fmt.Println(T[i].judul)
		}
			
	}
}