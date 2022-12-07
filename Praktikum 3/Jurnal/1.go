package main

import "fmt"

func main() {
	var num_ascii1, num_ascii2, num_ascii3, num_ascii4, num_ascii5 int
	var kata string
	fmt.Scanln(&num_ascii1, &num_ascii2, &num_ascii3, &num_ascii4, &num_ascii5)
	fmt.Scanln(&kata)
	fmt.Printf("%c", num_ascii1)
	fmt.Printf("%c", num_ascii2)
	fmt.Printf("%c", num_ascii3)
	fmt.Printf("%c", num_ascii4)
	fmt.Printf("%c", num_ascii5)
	fmt.Println()

	fmt.Printf("%c", kata[0]+1)
	fmt.Printf("%c", kata[1]+1)
	fmt.Printf("%c", kata[2]+1)
}