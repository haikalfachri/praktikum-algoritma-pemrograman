package main

import "fmt"

func main() {
	var kunciRahasia, a, b, i int
	fmt.Scanln(&kunciRahasia)
	i = 1
	for i != 0 {
		fmt.Scanln(&a, &b)
		if a%kunciRahasia == 0 {
			fmt.Println(b)
		} else if b%kunciRahasia == 0 {
			fmt.Println(a)
		} else if a == -1 && b == -1 {
			i = i - 1
		}
	}
}
