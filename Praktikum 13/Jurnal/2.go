package main
import "fmt"
const NMAX = 100
func main(){
	var n, d, u int
	var data [NMAX]int

	fmt.Scan(&n, &d, &u)
	isiArray(&data, &n)
	sorting(&data, u, d, n)
	tampilArray(data, n)

}
func isiArray(t *[100]int, n *int){
	/* IS. terdefinisi bilangan bulat n, dan n buah bilangan bulat telah siap pada
	piranti masukan
	FS. array t berisi n buah bilangan bulat yang berasal dari user */
	for i:=0;i<*n;i++{
		fmt.Scan(&t[i])
	}
}
func tampilArray(t [100]int, n int){
	/* IS. terdefinisi sebuah array t yang berisi n buah bilang bulat
	FS. menampilkan array t ke layar secara horizontal */
	for i:=0;i<n;i++{
		fmt.Printf("%d ",t[i])
	}
}
func sorting(t *[100]int, u,d,n int){
	/* IS. terdefinisi sebuah array t yang berisi n bilangan bulat, dan indeks bilangan
	bulat u dan d
	FS. array t terurut descending dari indeks ke-u hingga ke-d dengan menggunakan
	algoritma insertion sort */
	for i:=d+1;i<u;i++{
		for j:=d;j<u;j++{
			if t[j] < t[i]{
				temp := t[j]
				t[j] = t[i]
				t[i] = temp
			}
		}
	}
}