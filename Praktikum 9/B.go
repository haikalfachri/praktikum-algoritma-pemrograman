package main

import "fmt"

type tabel = [127]string

func isiArray(t *tabel, n *int){
	var masukkan string

	*n = 1
	fmt.Print("Teks : ")
	fmt.Scan(&masukkan)
	for (masukkan != ".") && (*n >= 1  && *n <= 127){
		t[*n] = masukkan
		*n++
		fmt.Scan(&masukkan)
	}
}

func cetakArray(t tabel, n int){
	var teks string

	for i:=1;i<len(t);i++{
		teks = teks + t[i]
	}
	fmt.Println(teks)
}

func balikanArray(t *tabel, n *int){
	for i, j := 1, len(t) - 1; i < j; i, j = i+1, j-1{
		t[i], t[j] = t[j], t[i]
	}
}

func palindrom(t tabel, n int)bool{
	var status bool
	var teksNormal, teksReverse string

	for i:=1;i<len(t);i++{
		teksNormal += t[i]
	}
	fmt.Print("Normal : ")
	cetakArray(t, n)
	balikanArray(&t, &n)
	for i:=1;i<len(t);i++{
		teksReverse += t[i]
	}
	fmt.Print("Reverse : ")
	cetakArray(t, n)
	status = teksNormal == teksReverse
	return status

}

func main(){
	var t tabel
	var n int
	var status bool

	isiArray(&t, &n)
	status = palindrom(t, n)
	fmt.Println("Palindrom ? ",status)
}