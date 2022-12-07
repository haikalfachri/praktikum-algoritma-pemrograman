package main

import "fmt"

func main() {
	var winner, player, temp string
	var A, B int
	var i, ronde int
	var nilai, answer int

	winner = "A"
	player = "B"
	ronde = 1
	fmt.Printf("Ronde", ronde, ":")
	fmt.Println(winner, "- masukkan angka rahasia:")
	fmt.Scanln(&nilai)
	for nilai != -101 {
		i = 1
		fmt.Println(player, "- masukkan angka tebakan ke-", i, ":")
		fmt.Scanf(&answer)
		for i < 3 && answer != nilai {
			i++
			fmt.Println(player, "- masukkan angka tebakan ke-", i, ":")
			fmt.Scanf(&answer) 
		} if nilai == answer {
			temp = winner
			winner = player
			player = temp 
		} fmt.Println(winner, "adalah pemenangnya")
		ronde++
		fmt.Println(winner, "masukkan angka rahasia: ")
		fmt.Scanf(&nilai) 
		} fmt.Println("Permainan selesai")
	}
