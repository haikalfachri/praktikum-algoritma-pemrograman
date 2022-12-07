package main

import "fmt"

func main() {
	var n, x, y, i int
	fmt.Scanln(&n)
	i = 0
	for i < n {
		fmt.Scanln(&x, &y)
		if (x+y)%2 == 0 {
			fmt.Println(x)
		} else if (x+y)%2 == 1 {
			fmt.Println(y)
		}
		i++
	}
}
