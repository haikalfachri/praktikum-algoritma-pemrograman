package main

import "fmt"

func findFactorial(n int, result *int) {
	*result = 1
	for n > 1 {
		*result = *result * n
		n--
	}
}

func permutation(n, r int) int {
	var x, y int
	findFactorial(n, &x)
	findFactorial((n - r), &y)
	return x / y
}

func combination(n, r int) int {
	var x, y, z int
	findFactorial(n, &x)
	findFactorial((n - r), &y)
	findFactorial(r, &z)
	return x / (y * z)
}

func main() {
	var n, r int
	fmt.Scan(&n, &r)
	fmt.Println(permutation(n, r), combination(n, r))
}
