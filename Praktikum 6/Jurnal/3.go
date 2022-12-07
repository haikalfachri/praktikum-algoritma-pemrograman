package main

import "fmt"
import "math"

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func posisi(cx, cy, r, x, y float64) bool {
	var pos float64
	pos = jarak(cx, cy, x, y)
	return pos < r
}

func main(){
	var cx, cy, r, x, y float64
	fmt.Scan(&cx, &cy, &r)
	fmt.Scan(&x, &y)
	if posisi(cx, cy, r, x, y){
		fmt.Println("Anda berada dalam taman")
	} else{
		fmt.Println("Anda berada di luar taman")
	}

}