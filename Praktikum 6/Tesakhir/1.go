package main

import "fmt"

func min(f1, f2 float64) float64 {
	if f1 > 0 && f2 > 0 {
		if f1 < f2 {
			return f1
		}else {
			return f2
		}
	}else if f1 > 0 && f2 <= 0 {
		return f1
	}
	return f2
}


func max(f1, f2 float64) float64 {
	if f1 > 0 && f2 > 0 {
		if f1 > f2 {
			return f1
		}else {
			return f2
		}
	}else if f1 > 0 && f2 <= 0 {
		return f1
	}
	return f2
}

func main() {
	var a, b, c, d, e, f, g, h, i, j float64
	var maks, mins float64
	var makx, minx float64

	fmt.Scanf("%v %v %v %v %v %v %v %v %v %v", &a, &b, &c, &d, &e, &f, &g, &h, &i, &j)
	mins = min(a, b)
	mins = min(mins, c)
	mins = min(mins, d)
	mins = min(mins, e)
	mins = min(mins, f)
	mins = min(mins, g)
	mins = min(mins, h)
	mins = min(mins, i)
	mins = min(mins, j)
	minx = mins

	maks = max(a, b)
	maks = max(maks, c)
	maks = max(maks, d)
	maks = max(maks, e)
	maks = max(maks, f)
	maks = max(maks, g)
	maks = max(maks, h)
	maks = max(maks, i)
	maks = max(maks, j)
	makx = maks

	fmt.Print(makx, minx)

}