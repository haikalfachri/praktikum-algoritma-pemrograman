package main

import "fmt"

type Buku struct{
	judul string
	penulis string
	tahun int
}
type tabBuku = [5]Buku

func tambahBuku(kardus *tabBuku, atas *int){
	var tambah Buku

	fmt.Scan(&tambah.judul, &tambah.penulis, &tambah.tahun)
	*atas++
	kardus[*atas] = tambah
}

func ambilBuku(kardus *tabBuku, atas *int, ambil *Buku){
	*ambil = kardus[*atas]
	*atas--
}
func cariBuku(kardus *tabBuku, atas *int, x string){
	var ambil Buku

	for *atas>=0{
		ambilBuku(&*kardus, &*atas, &ambil)
		if ambil.judul == x{
			fmt.Println("Ketemu")
			break
		}else if ambil.judul != x{
			if *atas >= 0{
				fmt.Println(ambil.judul)
			}else{
				fmt.Println(ambil.judul)
				fmt.Println("Tidak Ketemu")
			}
		}
	}
}
func main(){
	var kardus tabBuku 
	var atas, jumlah int
	var x string

	atas = -1	//{1}
	fmt.Print("Masukkan jumlah buku yang akan dimasukkan ke dalam kardus? (maks 5 buku) : ")
	fmt.Scan(&jumlah)
	fmt.Printf("Masukkan %v buku tersebut ke dalam kardus\n(Masukkan berupa judul, penulis, tahun terbit dipisah dengan spasi) : ",jumlah)
	for i:=1; i<=jumlah; i++{
		tambahBuku(&kardus, &atas)//{2}
	}
	fmt.Print("Cari buku apa? ")
	fmt.Scan(&x)
	cariBuku(&kardus, &atas, x)//{3}
	if atas >= 0 && atas < 4{
		butuh := 4-atas
		fmt.Print("Anda butuh ", butuh, " buku supaya kardus penuh\n(Masukkan berupa judul, penulis, tahun terbit dipisah dengan spasi) : ")
		for i:=1; i<=butuh; i++{
			tambahBuku(&kardus, &atas)//{4}
		}
	}
	fmt.Print("Cari buku apa? ")
	fmt.Scan(&x)
	cariBuku(&kardus, &atas, x)//{5}
}