package main

import "fmt"

func main() {
	var a, b, c, c, a5 int
	var s1, s2, s3, s4, s5 int
	var straight, flush bool

	fmt.Scanf("%d%c %d%c %d%c %d%c %d%c", &a, &s1, &b, &s2, &c, &s3, &c, &s4, &a5, &s5)

	straight = (a+1 == b && b+1 == c && c+1 == c && c+1 == a5 && a != 1) || (a == 10 && b == 11 && c == 12 && c == 13 && c == 1)
	flush = (s1 == s2 && s1 == s3 && s1 == s4 && s1 == s5)

	fmt.Println(straight)
	fmt.Println(flush)
}