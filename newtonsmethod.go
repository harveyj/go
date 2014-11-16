package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    z:=100.0
    nz:=0.0
    for nz-z > 0.0001 || z-nz > 0.0001 {
        nz=z
        z=z-(z*z-x)/(2*z)
    }
    return z
}

func main() {
    a:=Sqrt(300000)
    fmt.Println(a, a*a)
}