package main

import "fmt"

func hitungSuara(suara int, calon1, calon2, calon3 *int){
	if suara == 1{
		*calon1++
	}else if suara == 2{
		*calon2++
	}else if suara == 3{
		*calon3++
	}
}

func main(){
	var suara int
	var calon1, calon2, calon3 int

	calon1, calon2, calon3 = 0, 0, 0
	fmt.Scan(&suara)
	for suara != -1{
		hitungSuara(suara, &calon1, &calon2, &calon3)
		fmt.Scan(&suara)
	}
	fmt.Printf("%d %d %d", calon1, calon2, calon3)
}