package main

import "fmt"

func main(){
	var masukkan int
	var array [21]int
	var masuk, sah int = 0, 0
	var ketua, wakil int = 0, 0

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
	fmt.Println(array)
	fmt.Println("Suara masuk: ",masuk)
	fmt.Println("Suara sah: ",sah)
	for i:=0;i<len(array);i++{
		if array[i] != 0{
			if array[i] >= array[ketua]{
				if array[i] > array[ketua]{
					wakil = ketua
					ketua = i
				}else if array[i] == array[ketua]{
					wakil = i
				}
			}else{
				if array[i] < wakil{
					wakil = i
				}
			}
		}
	}
	fmt.Println("Ketua RT: ",ketua)
	fmt.Println("Wakil ketua: ",wakil)
}