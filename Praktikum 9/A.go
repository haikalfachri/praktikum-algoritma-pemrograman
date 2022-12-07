package main

import "fmt"

type tabGol = [50]int

func inputData(t *tabGol, n *int){
	var masukkan int

	fmt.Scan(&masukkan)
	for masukkan >= 0 && *n <= 50{
		t[*n] = masukkan
		fmt.Scan(&masukkan)
		*n++
	}
}

func rataan(t tabGol, n int)float64{
	var total int
	var rata float64

	total = 0
	for i:=0;i<n;i++{
		total += t[i]
	}
	 rata = float64(total) / float64(n)
	 return rata
}

func main(){
	var tim1, tim2, tim3 float64
	var t tabGol
	var n int

	n = 0
	for i:=1;i<=3;i++{
		fmt.Printf("Tim ke-%d : ", i)
		inputData(&t, &n)
		if i == 1{
			tim1 = rataan(t, n)
		}else if i == 2{
			tim2 = rataan(t, n)
		}else if i == 3{
			tim3 = rataan(t, n)
		}
		for i:=0;i<n;i++{
			t[n] = 0
		}
		n = 0
	}
	fmt.Printf("Tim 1 : %v , Tim 2 : %v , Tim 3 : %v",tim1,tim2,tim3)



}