package main
import "fmt"
import "math/rand"
func main() {
	 var seed int64
	 var jenis, ukuran string
 	 fmt.Print("Angka rahasia: ")
 	 fmt.Scanln(&seed)
 	 rand.Seed(seed)
 	 fmt.Println("Satu nilai 1...6:", rand.Intn(6)+1)
}

