package main

import "fmt"

func main(){
	data := make([]string, 0)

	var input string

	fmt.Scan(&input)
	for input != "."{
		data = append(data, input)
		fmt.Scan(&input)
	}

	jumlah := make(map[string]int)
	
	for _, trend := range data{		
		_, i := jumlah[trend]
		if i {
			jumlah[trend] += 1 
		} else {
			jumlah[trend] = 1 
	}
	var count int = 0
	var trending string
	for k, v := range jumlah {
		if v > count{
			trending = k
			count = v
		}
	}
	fmt.Println(trending)
	
}