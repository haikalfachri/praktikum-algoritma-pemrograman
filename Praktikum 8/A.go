package main

import "fmt"

func main(){
	var p, q, n int
	var sisaP, sisaQ int

	fmt.Scan(&n, &p, &q)
	if p == q{
		fmt.Print("Kita tidak pernah bertemu")	
	}else if n % p == 0 && n % q != 0{
		fmt.Print("Kita akan bertemu hari ini")
	}else if n % q == 0 && n % p != 0{
		fmt.Print("Kita akan bertemu hari ini")
	}else if n % p == 0 && n % q == 0{
		if p > q{
			fmt.Printf("Kita akan bertemu %d hari lagi", q)
		}else{
			fmt.Printf("Kita akan bertemu %d hari lagi", p)
		}
	}else{
		sisaP = p - (n % p)
		sisaQ = q - (n % q)
		if sisaP > sisaQ{
			fmt.Printf("Kita akan bertemu %d hari lagi", sisaQ)
		}else{
			fmt.Printf("Kita akan bertemu %d hari lagi", sisaP)
		}
	}
}