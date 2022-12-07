t "fmt"

func main() {
	var bil, i, j int
	fmt.Scanln(&bil)
	j = 0
	for i = 1; i <= bil; i++ {
		if bil%i == 0 {

			fmt.Print(i, " ")
			j++
		}
	}
	if j == 2 {
		fmt.Println("\nOleole")

	} else {
		fmt.Println("\nBukan Oleole")
	}
}