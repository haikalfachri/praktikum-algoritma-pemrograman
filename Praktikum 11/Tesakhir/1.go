package main

import "fmt"

func cariMaxMin(array []int) (min int, max int) {
	min = array[0]
	max = array[0]
	for _, bil := range array {
		if bil < min {
			min = bil
		}
		if bil > max {
			max = bil
		}
	}
	return min, max
}

func main(){
	var n, cariBil, bil int
	var posisiMin, posisiMax int
	var max, min int
	var ketemu bool

	fmt.Scan(&n, &cariBil)
	dataBil := make([]int, 0)
	for i:=0;i<n;i++{
		fmt.Scan(&bil)
		dataBil = append(dataBil, bil)
	}
	min, max = cariMaxMin(dataBil)
	for i:=0;i<n;i++{
		if dataBil[i] == min{
			posisiMin = i
		}
		if dataBil[i] == max{
			posisiMax = i
		}
	}
	if posisiMax < posisiMin{
		temp := posisiMax
		posisiMax = posisiMin
		posisiMin = temp
	}
	if cariBil % 2 == 0{
		dataBilx := dataBil[posisiMin:posisiMax+1]
		fmt.Println(dataBilx)
		for i:=0;i<len(dataBilx);i++{
			if dataBilx[i] == cariBil{
				ketemu = true
				break
			}
		}
		if ketemu{
			fmt.Println("ditemukan")
		}else{
			fmt.Println("tidak ditemukan")
		}
	}else if cariBil % 2 != 0{
		dataBilx := dataBil[0:posisiMin+1]
		dataBily := dataBil[posisiMax:]
		fmt.Println(dataBilx)
		fmt.Println(dataBily)
		for i:=0;i<len(dataBilx);i++{
			if dataBilx[i] == cariBil{
				ketemu = true
				break
			}
		}
		for i:=0;i<len(dataBily);i++{
			if dataBily[i] == cariBil{
				ketemu = true
				break
			}
		}
		if ketemu{
			fmt.Println("ditemukan")
		}else{
			fmt.Println("tidak ditemukan")
		}
	}
}