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
	var temp int
	// lakukan proses input suara secara berulang di sini, simpan ke dalam array p, sehingga terdapat array p yang berisi hasil peroleh suara n partai.
	fmt.Scan(&input)
	for input != -1{
		vote = append(vote, input)
		fmt.Scan(&input)
	}
	for i:=0; i<len(vote); i++{
		if vote[i] != temp{
			p[i].nama = vote[i]
		}
		temp = vote[i]
	}
	for i:=0;i<len(vote);i++{
		j:=i
		for j<len(p){
			if vote[i] == p[j].nama {
				p[i].total += 1
			}
			j++
		}
		
	}
	fmt.Println()
	for i:=0;i<len(p);i++{
		if p[i].nama != 0{
			fmt.Println(p[i])
		}
		
	}
	// lakukan proses pengurutan dengan insertion sort descending berdasarkan jumlah suara yang diperoleh
	var n = len(p)
    for i := 1; i < n; i++ {
        j := i
        for j > 0 {
            if p[j-1].total > p[j].total {
                p[j-1], p[j] = p[j], p[j-1]
            }
            j = j - 1
		}
	}
	// tampilkan array p
	fmt.Println()
	for i:=0;i<len(p);i++{
		if p[i].nama != 0{
			fmt.Println(p[i])
		}
		
	}
}