package main

import "fmt"

func f(x float64) float64 {
	return x * x
}

func g(x float64) float64 {
	return x - 2
}

func h(x float64) float64 {
	return x + 1
}

func komposisi(x float64, y *float64){
	*y = f(g(h(x))) 
}

func main(){
	var x, y float64
	fmt.Scan(&x)
	fmt.Printf("f(%v): %v\n", x, f(x))
	fmt.Printf("g(%v): %v\n", x, g(x))
	fmt.Printf("h(%v): %v\n", x, h(x))
	komposisi(x,&y)
	fmt.Println(y)
}