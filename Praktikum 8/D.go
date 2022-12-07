package main

import "fmt"

func hitungNopol(noPol string, nBandung, nJakarta, nLainnya *int ){
	if noPol == "B" || noPol == "F"{
		*nJakarta++
	}else if noPol == "D" || noPol == "Z"{
		*nBandung++
	}else{
		*nLainnya++
	}
}

func main(){
	var noPol string
	var nBandung, nJakarta, nLainnya int

	nBandung, nJakarta, nLainnya = 0, 0, 0
	fmt.Scan(&noPol)
	for noPol != "." {
		hitungNopol(noPol, &nBandung, &nJakarta, &nLainnya)
		fmt.Scan(&noPol)
	}
	fmt.Printf("%d %d %d", nBandung, nJakarta, nLainnya)
}