package main
import "fmt"
const NMAX = 100
type belanja [NMAX]int
func main(){
	// Deklarasi Variabel
	var t belanja
	var n, ha int
	var hp float64
	// Lakukan proses masukan
	entri(&t, &n)
	// Hitung total belanja
	ha = total(t, n)
	// tentukan apakah mendapatkan promo atau tidak
	if ha > 100000{
		// lakukan pengurutan
		urut(&t, n)
		// lakukan perhitungan promo
		promo(t, n, &hp)
		// tampilkan keluaran yang diminta
		fmt.Println(ha,hp)
	}else{
		// tampilkan keluaran yang diminta
		fmt.Println(ha,ha)
	}
}
func entri(t *belanja, n *int){
	/* IS. data belanja telah siap pada piranti masukan
	FS. array t berisi sejumlah n harga barang yang dibeli */
	var h, m int
	fmt.Scan(&h, &m)
	for h != 0 || m != 0{
		t[*n] = (h * m)
		*n++
		fmt.Scan(&h, &m)
	}
}
func urut(t *belanja, n int){
	/* IS. terdefinisi array t yang berisi n harga barang yang dibeli
	FS. array t terurut secara ascending berdasarkan harga barang dengan
	selection/insertion sort */
	for i:=1;i<n;i++{
		for j:=0;j<n;j++{
			if t[j] > t[i]{
				temp := t[j]
				t[j] = t[i]
				t[i] = temp
			}
		}
	}
}
func total(t belanja, n int) int {
	/* IS. terdefinisi array t yang berisi n harga barang yang dibeli
	FS. mengembalikan total harga barang yang terdapat pada array t */
	var totalHarga int
	for i:=0;i<n;i++{
		totalHarga += t[i]
	}
	return totalHarga
}
func promo(t belanja, n int, hp* float64){
	/* IS. terdefinisi array t yang berisi n harga barang yang dibeli dan terurut
	ascending
	FS. hp berisi total harga setiap barang setelah dikurangi dengan diskonnya*/
	j := 1
	for i:=0;i<n;i++{
		*hp += float64(t[i] - (t[i] * j / 100))
		j++
	}
}