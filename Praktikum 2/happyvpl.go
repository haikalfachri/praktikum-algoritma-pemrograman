package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	if (33 <= a && a <= 126) && a < b && (33 <= b && b <= 126) {
		t_a := a
		for t_a != b+1 {
			character := string(t_a)
			fmt.Printf("Simbol ASCII dari %d adalah %s\n", t_a, character)
			t_a++
		}
	}
}