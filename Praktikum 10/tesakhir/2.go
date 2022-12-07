package main
import "fmt"
const NMAX = 1000000
var data [NMAX]int
func main(){
	/* buatlah kode utama yang membaca baris pertama (n dan k). kemudian data diisi
	oleh prosedur isiArray(n), dan pencarian oleh fungsi posisi(n,k), dan setelah itu
	output dicetak.*/
	var n, k int
	var out int

	fmt.Scan(&n, &k)
	isiArray(n)
	out = posisi(n, k)
	if out == -1{
		fmt.Print("TIDAK ADA")
	}else{
		fmt.Print(out)
	}
}
func isiArray(n int){
	/* IS. Data n sudah siap pada piranti masukan.
	FS. Array data berisi n (<=NMAX) bilangan */
	var masukkan int

	for i:=0; i<n; i++{
		fmt.Scan(&masukkan)
		data[i] = masukkan
	}
}
func posisi(n, k int) int {
	/* mengembalikan posisi k dalam array data dengan n elemen. Posisi dimulai dari
	posisi 0. Jika tidak ada kembalikan -1 */
	for i:=0; i<n; i++{
		if data[i] == k{
			return i
		}
	}
	return -1
}