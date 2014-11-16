package main

import "fmt"

func Cbrt(x complex128) complex128 {
    z := complex(100.0, 100.0)
    for i := 0; i < 100; i++ {
        z = z - (z*z*z - x) / (3*z*z)
    }
    return z
}

func main() {
    a := Cbrt(2)
    fmt.Println(a*a*a)
}