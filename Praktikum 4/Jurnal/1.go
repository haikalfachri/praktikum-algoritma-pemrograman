package main

import "fmt"


func main() {
	var n, i, m, h1, j int
	fmt.Scanln(&n)
	i = 0
	h1 = 0
	for i < n {
		fmt.Scanln(&inputan)
		if inputan > -1 && inputan < 10 {
			j = 0
			for j < i {
				m = m * 10
				j++
			}
			h1 = h1 + m
			i++
		}
	}
	fmt.Println(h1)

}
