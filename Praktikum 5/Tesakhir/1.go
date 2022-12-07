package main

import "fmt"

func main() {
	var bil1, bil2 int
	var netral, positive int

	netral = 0
	positive = 0
	
	fmt.Scan(&bil1, &bil2)
	if bil1 + bil2 == 0 {
		netral++
	}else if bil1 + bil2 > 0{
		positive++
	}
	for bil1 + bil2 >= 0{
		fmt.Scan(&bil1, &bil2)
		if bil1 > bil2{
			if bil1 + bil2 == 0{
				netral++
			}else if bil2 + bil1 > 0{
				positive++
			}
		}else{
			if bil2 + bil1 == 0{
				netral++
			}else if bil2 + bil1 > 0{
				positive++
			}
		}
	}
	fmt.Println("Netral: ", netral)
	fmt.Println("Positive: ", positive)
}

	