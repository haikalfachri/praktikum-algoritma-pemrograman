package main

import "fmt"

func main() {
	var S string
	var A, B int
	var hasil_penjumlahan int
	fmt.Scanln(&S, &A, &B)
	hasil_penjumlahan = A + B
	fmt.Println("Kata =", S)
	fmt.Println("Jumlah =", hasil_penjumlahan)

}
