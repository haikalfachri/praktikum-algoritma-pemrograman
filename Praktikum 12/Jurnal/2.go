package main

import "fmt"
import "sort"

func main(){
	var input int
	
	dataPil := make([]int,0)
	fmt.Scan(&input)
	for input != -1{
		dataPil = append(dataPil, input)
		fmt.Scan(&input)
	}
	dataPilx := make(map[int]int)
	for _, pil := range dataPil{
		_, i := dataPilx[pil]
		if i{
			dataPilx[pil] += 1
		}else{
			dataPilx[pil] = 1
		}
	}
    keys := make([]int, 0, len(dataPilx))
    for key := range dataPilx {
        keys = append(keys, key)
    }

    sort.Slice(keys, func(i, j int) bool {
        return dataPilx[keys[i]] > dataPilx[keys[j]]
    })

	for k, v := range dataPilx{
		fmt.Printf("%v(%v) ",k, v)
	}
}