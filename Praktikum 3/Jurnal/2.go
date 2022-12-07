package main

import "fmt"

func main() {
	var benar int
	var tebakan1, tebakan2 int
	var tebakTebakan int
	var a1, a2, a3, a4 int
	var b1, b2, b3 int

	fmt.Scanln(&benar)
	fmt.Scanln(&tebakan1)
	fmt.Scanln(&tebakan2)
	fmt.Scanln(&tebakTebakan)

	a1 = benar / 1000
	a2 = benar / 100 % 10
	a3 = benar / 10 % 10
	a4 = benar % 10

	b1 = benar / 100
	b2 = benar / 1000 % 100
	b3 = benar % 100

	fmt.Println(tebakTebakan == benar)
	fmt.Println(tebakan1 == a1 || tebakan1 == a2 || tebakan1 == a3 || tebakan1 == a4)
	fmt.Println(tebakan2 == b1 || tebakan2 == b2 || tebakan2 == b3)

}
