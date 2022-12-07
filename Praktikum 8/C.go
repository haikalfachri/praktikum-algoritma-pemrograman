package main

import "fmt"

func main(){
	var M, N int
	var kpk int

	kpk = 1
	fmt.Scan(&M, &N)
	for{
		if kpk % M == 0 && kpk % N == 0{
			fmt.Print(kpk)
			break
		}
		kpk++
	}
}