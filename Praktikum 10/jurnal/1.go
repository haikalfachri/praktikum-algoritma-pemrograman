package main

import "fmt"

func main(){
	var masukkan int
	var array [21]int
	var masuk, sah int

	fmt.Scan(&masukkan)
	for masukkan != 0{
		if masukkan >= 1 && masukkan <=20{
			array[masukkan]++
			masuk++
			sah++
		}else{
			masuk++
		}
		fmt.Scan(&masukkan)
	}
	fmt.Println("Suara masuk: ",masuk)
	fmt.Println("Suara sah: ",sah)
	for i:=1;i<len(array);i++{
		if array[i] != 0{
			fmt.Println(i, ":", array[i])
		}
	}
}