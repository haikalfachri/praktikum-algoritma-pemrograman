package main

import "fmt"

func main() {
	var n, nilai, besar int
	iterasi := 1
	banyak := 1

	fmt.Scan(&n)
	for iterasi <= n {
		fmt.Scan(&nilai)

		if nilai > besar {
			besar = nilai
			banyak = 1

		} else if nilai == besar {
			banyak++

		}
		iterasi++
	}
	fmt.Printf("Nilai terbesar adalah %d yang diperoleh %d orang mahasiswa.", besar, banyak)
}
