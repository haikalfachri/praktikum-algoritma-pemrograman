package main

import "fmt"

const NMAX = 1001
type info struct{
	namaDepan string
	namaBelakang string
	gol int
	assist int
}
type tab [NMAX]info

func main(){
	var n int
	var data tab

	fmt.Scan(&n)
	for i:=0;i<n;i++{
		fmt.Scan(&data[i].namaDepan, &data[i].namaBelakang, &data[i].gol, &data[i].assist)
	}
	for i := 0; i < n-1; i++ {
        min := i
        for j := i + 1; j < n; j++ {
            if data[j].gol >= data[min].gol {
                if data[j].gol > data[min].gol {
					min = j
				}else{
					if data[j].assist > data[min].assist {
						min = j
					}
				}
            }
        }
        data[i], data[min] = data[min], data[i]
    }
	fmt.Println()
	for i:=0;i<n;i++{
		fmt.Println(data[i].namaDepan, data[i].namaBelakang, data[i].gol, data[i].assist)
	}
}