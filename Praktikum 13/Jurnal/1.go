package main
import "fmt"
const NMAX = 100
type mahasiswa struct{
	inisial string
	tinggi int
}
type dataMhs [NMAX]mahasiswa
func main(){
	var n int
	var data dataMhs

	fmt.Scan(&n)
	bacaData(&data, &n)
	urutData(&data, n)
	fmt.Println()
	cetakData(data, n)
}
func bacaData(t *dataMhs, n *int){
	/* IS. n data mahasiswa telah siap pada piranti masukan
	FS. menerima input n dan array t berisi n data tinggi mahasiswa */
	for i:=0;i<*n;i++{
		fmt.Scan(&t[i].inisial, &t[i].tinggi)
	}
}	
func cetakData(t dataMhs, n int){
	/* IS. terdefinisi sebuah array t yang berisi n data tinggi mahasiswa
	FS. menampilkan array t ke layar monitor */
	for i:=0;i<n;i++{
		fmt.Println(t[i].inisial, t[i].tinggi)
	}
}
func urutData(t *dataMhs, n int){
	/* IS. terdefinisi sebuah array t yang berisi n data tinggi mahasiswa
	FS. array t terurut ascending berdasarkan tinggi dengan algoritma selection sort
	*/
	var tempTinggi mahasiswa

	for i:=0;i<n;i++{
		temp := i
		for j:=temp+1;j<n;j++{
			if t[temp].tinggi < t[j].tinggi{
				temp = j
			}
		}
		tempTinggi = t[i]
		t[i] = t[temp]
		t[temp] = tempTinggi
	}
}