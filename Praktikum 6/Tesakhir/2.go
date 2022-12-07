package main

import "fmt"

func f(x float64) float64 { // mencari nilai f(x)
	return (2 * x) + 5
}
func g(x float64) float64 { // mencari nilai g(x)
	return (x * x) + (2 * x)
}
func h(x float64) float64 { // mencari nilai h(x)
	return x - 3
}
func fgh(x float64, y *float64) { // mencari nilai (fogoh)(x)
	*y = f(g(h(x)))
}
func fhg(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = f(h(g(x)))
}
func fhf(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = f(h(f(x)))
}
func fgf(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = f(g(f(x)))
}
func ffh(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = f(f(h(x)))
}
func ffg(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = f(f(g(x)))
}
func fgg(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = f(g(g(x)))
}
func fhh(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = f(h(h(x)))
}
func fff(x float64, y *float64) { // mencari nilai (fofof)(x)
	*y = f(f(f(x)))
}
func ghf(x float64, y *float64) { // mencari nilai (gohof)(x)
	*y = g(h(f(x)))
}
func gfh(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = g(f(h(x)))
}
func gfg(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = g(f(g(x)))
}
func ghg(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = g(h(g(x)))
}
func ggf(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = g(g(f(x)))
}
func ggh(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = g(g(h(x)))
}
func gff(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = g(f(f(x)))
}
func ghh(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = g(h(h(x)))
}
func ggg(x float64, y *float64) { // mencari nilai (gogog)(x)
	*y = g(g(g(x)))
}
func hfg(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = h(f(g(x)))
}
func hgf(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = h(g(f(x)))
}
func hgh(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = h(g(h(x)))
}
func hfh(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = h(f(h(x)))
}
func hhg(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = h(h(g(x)))
}
func hhf(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = h(h(f(x)))
}
func hff(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = h(f(f(x)))
}
func hgg(x float64, y *float64) { // mencari nilai (hofog)(x)
	*y = h(g(g(x)))
}
func hhh(x float64, y *float64) { // mencari nilai (hohoh)(x)
	*y = h(h(h(x)))
}

func main() {
	var x, y float64
	var a, b, c string
	fmt.Scanln(&x, &a, &b, &c) 
	if a == "f" && b == "g" && c == "h" {
		hgf(x, &y)
	} else if a == "g" && b == "f" && c == "h" {
		hfg(x, &y)
	} else if a == "h" && b == "g" && c == "h" {
		hgh(x, &y)
	} else if a == "h" && b == "f" && c == "h" {
		hfh(x, &y)
	} else if a == "f" && b == "h" && c == "h" {
		hhf(x, &y)
	} else if a == "g" && b == "h" && c == "h" {
		hhg(x, &y)
	} else if a == "h" && b == "h" && c == "h" {
		hhh(x, &y)
	} else if a == "h" && b == "f" && c == "g" {
		gfh(x, &y)
	} else if a == "f" && b == "h" && c == "g" {
		ghf(x, &y)
	} else if a == "h" && b == "g" && c == "g" {
		ggh(x, &y)
	} else if a == "f" && b == "g" && c == "g" {
		ggf(x, &y)
	} else if a == "g" && b == "f" && c == "g" {
		gfg(x, &y)
	} else if a == "g" && b == "h" && c == "g" {
		ghg(x, &y)
	} else if a == "g" && b == "g" && c == "g" {
		ggg(x, &y)
	} else if a == "h" && b == "h" && c == "g" {
		ghh(x, &y)
	} else if a == "f" && b == "f" && c == "g" {
		gff(x, &y)
	} else if a == "f" && b == "f" && c == "h" {
		hff(x, &y)
	} else if a == "g" && b == "h" && c == "f" {
		fhg(x, &y)
	} else if a == "h" && b == "g" && c == "f" {
		fgh(x, &y)
	} else if a == "f" && b == "g" && c == "f" {
		fgf(x, &y)
	} else if a == "f" && b == "h" && c == "f" {
		fhf(x, &y)
	} else if a == "g" && b == "g" && c == "f" {
		fgg(x, &y)
	} else if a == "h" && b == "h" && c == "f" {
		fhh(x, &y)
	} else if a == "g" && b == "f" && c == "f" {
		ffg(x, &y)
	} else if a == "h" && b == "f" && c == "f" {
		ffh(x, &y)
	} else if a == "f" && b == "f" && c == "f" {
		fff(x, &y)
	} else if a == "g" && b == "g" && c == "h" {
		hgg(x, &y)
	} 
	fmt.Println(y)
}