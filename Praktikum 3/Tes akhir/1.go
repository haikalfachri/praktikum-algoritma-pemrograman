package main

import "fmt"

func main() {
	var a, b, c, d, e byte
	var geser int
	var a1, b1, c1, d1, e1 int

	fmt.Scanf("%c%c%c%c%c\n", &a, &b, &c, &d, &e)
	fmt.Scanf("%d", &geser)

	a1 = int(a) + geser
	b1 = int(b) + geser
	c1 = int(c) + geser
	d1 = int(d) + geser
	e1 = int(e) + geser

	fmt.Printf("%c%c%c%c%c", a1, b1, c1, d1, e1)
}