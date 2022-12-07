package main
import "fmt"
const NMAX = 1000000
// tipe bentukan partai
	type tabPartai struct{
		nama int
		total int
	}
// tipe array of partai dengan kapasitas NMAX
	type array [NMAX]tabPartai
func main(){
	// deklarasi variabel
	var input int
	var vote []int
	var p array
	var n int
	// lakukan proses input suara secara berulang di sini, simpan ke dalam array p, sehingga terdapat array p yang berisi hasil peroleh suara n partai.
	fmt.Scan(&input)
	for input != -1{
		vote = append(vote, input)
		fmt.Scan(&input)
	}
	for i:=0;i<len(vote);i++{
		for j:=0;j<len(p);j++{
			if vote[i] == j{
				p[j].nama = j	// [1 1 2 3 4 1]
				p[j].total += 1 // [{0 0} {0 0} {0 0} {0 0}]
				n++
			}	
		}
		
	}
	// lakukan proses pengurutan dengan insertion sort descending berdasarkan jumlah suara yang diperoleh
	for i:=0; i<n; i++ {
        min := i
        for j:=i+1; j<n; j++ {
            if p[j].total > p[min].total {
                min = j
            }else if p[j].total == p[min].total{
				if p[j].nama > p[min].nama{
					min = j
				}
			}
        }
        p[i], p[min] = p[min], p[i]
    }

	// tampilkan array p
	fmt.Println()
	for i:=0;i<n;i++{
		if p[i].total != 0{
			fmt.Printf("%v(%v) ",p[i].nama, p[i].total)
		}
	}
}