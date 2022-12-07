package main

import "fmt"

func main() {
    var r, t float64

    fmt.Scanln(&r, &t)
    fmt.Println(hitungVolume(r, t))
}

func hitungVolume(r float64, t float64) float64 {
    pi := 3.14
    vol := r * r * pi * t
    return vol
}