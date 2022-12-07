package main

import "fmt"

func main() {
	var nomor int
	fmt.Scanln(&nomor)
	if nomor%3 == 0 {
		fmt.Println("Fizz")
	}
	if nomor%5 == 0 {
		fmt.Println("Bazz")
	}
}