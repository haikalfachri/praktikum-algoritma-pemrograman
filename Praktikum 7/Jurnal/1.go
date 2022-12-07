package main

import "fmt"

func Num2Str(x int) string {
	if x == 0 {
		return "0"
	}else if x == 1 {
		return "1"
	}
	return ""
}

func Division(a int, b int, result *int, remainder *int){
	*result = a / b
	*remainder = a % b
}

func Des2Bin(desimal int) string {
	var hasil string
	var rDiv, rMod int

	hasil = ""
	for desimal > 0{
		Division(desimal, 2, &rDiv, &rMod)
		desimal = rDiv
		hasil = Num2Str(rMod) + hasil
	}
	return hasil
}

func main(){
	var biner string
	var desimal int

	fmt.Scan(&desimal)
	biner = Des2Bin(desimal)
	fmt.Print(biner)
}