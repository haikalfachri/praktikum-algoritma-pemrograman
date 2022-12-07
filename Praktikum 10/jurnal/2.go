package main

import "fmt"

func posisi(tab [107]string, n int, x string) int {
	/*mengembalikan indeks dari x pada array tab, apabila x ditemukan di dalam tab yang 
	berisi n buah nilai, -1 apabila tidak ditemukan*/
	for i:=0; i<n; i++{
		if tab[i] == x{
			return i
		}
	}
	return -1
}

func hapus(tab *[107]string, n *int, x string){
	/*IS. terdefinisi sebuah array tab yang berisi n buah nim mahasiswa, dan x adalah nim 
	mahasiswa yang wisuda
	FS. apabila x ditemukan pada tab, maka nim dihapus, seluruh elemen setelahnya 
	bergeser, dan n berkurang 1. tab dan n tidak berubah apabila x tidak di temukan*/
	var indeks, akhir int

	for i:=0; i<*n; i++{
		indeks = posisi(*tab, *n, x)
		if tab[indeks] == x{
			akhir = *n
			for j:=i; j<akhir; j++{
				tab[j] = tab[j+1]
				tab[akhir] = ""
				akhir--
			}
			*n--
		}
	}
}

func updateKelulusan(mhs *[107]string, wisuda [107]string, n *int, m int){
	/* IS. terdefinisi array mhs dan wisuda yang berisi sejumlah n dan m nim mahasiswa
	FS. seluruh nim mahasiswa wisuda dihapus dari array mhs, nilai n terupdate */
	var indeks int

	if *mhs == wisuda{
		for i:=0;i<*n;i++{
			mhs[i] = "0"
		}
	}else{
		for i:=0; i<len(wisuda); i++{
			x := wisuda[i]
			indeks = posisi(*mhs, *n, x)
			if indeks != -1{
				hapus(&*mhs, &*n, x)
			}
		}
	}
}


func main(){
	var n, m int
	var daftar, wisuda [107]string
	var x string

// lakukan proses input untuk baris pertama
	fmt.Scan(&n)
	for i:=0; i<n; i++{
		fmt.Scan(&x)
		daftar[i] = x
	}
// lakukan proses input untuk baris kedua
	fmt.Scan(&m)
	for i:=0; i<m; i++{
		fmt.Scan(&x)
		wisuda[i] = x
	}
// panggil prosedur update kelulusan di sini
	updateKelulusan(&daftar, wisuda, &n, m)
// lakukan proses output array daftar di sini
	fmt.Print(n)
	for i:=0;i<n;i++{
		fmt.Printf(" %v", daftar[i])
	}
}