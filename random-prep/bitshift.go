package main

import (
    "fmt"
)

func main(){
    var a byte = 0xA2
    b := a
    fmt.Printf("initial: %08b\n", a)
    fmt.Printf("a>>4: %08b\n", a>>4)
    fmt.Printf("b<<4: %08b\n", b<<4)
    b = a>>4 + b<<4
    fmt.Printf("a + b: %08b\n", b)
}