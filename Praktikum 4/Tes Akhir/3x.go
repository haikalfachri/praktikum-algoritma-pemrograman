package main
import "fmt"
import "math/rand"
func main() {
	var seed int64
	var anda, dadang, nilai, angka int

	fmt.Scanln(&seed)
	rand.Seed(seed)
	dadang = rand.Intn(6)+1
	nilai = rand.Intn(6)+1
	fmt.Print("Angka Rahasia: ")
	fmt.Scan(&angka)
	fmt.Print("Anda: ")
	fmt.Scan(&anda)
	fmt.Print("Dadang: %d", dadang)
	if anda == nilai && dadang == nilai {
		fmt.Print("Nilai dadu %d, Seri", nilai)
	} else if anda == nilai && dadang != nilai {
		fmt.Print("Nilai dadu %d, Pemenang adalah anda", nilai)
	} else if anda != nilai && dadang == nilai {
		fmt.Print("Nilai dadu %d, Pemenang adalah Dadang", nilai)
	} else {
		fmt.Print("Nilai dadu %d, Tidak ada pemenang", nilai)
	}


}
