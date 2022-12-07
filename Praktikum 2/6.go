package main

import "fmt"

func main() {
    var nilai int
    var mean float64
    jumlah := 0
    n := 0

    fmt.Scan(&nilai)
    for int(nilai) != -1 {
        n += 1
        jumlah += nilai
        fmt.Scan(&nilai)
    }
    if n == 0 {
        mean = 0.0
    } else {
        mean = float64(jumlah) / float64(n)
    }

    fmt.Println(mean)
}