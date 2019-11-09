package main

import (
  //  "fmt"
    "github.com/01-edu/z01"
)
func main() {

    number := 153212345
    n := number
    minus := false
    if number < 0 {
        minus = true
        number *= -1 
    }
    
    len := 0
    for n !=0 {
        n /= 10
        len ++
    }

    if minus == true {
        len = len + 1
    }

    array := make([]rune, len)
    for i := len-1; i >= 0; i-- {
        array[i] = rune(number%10) + '0'
        number /= 10
    }
    if minus == true {
        array[0] = '-'
    }
    for _,s := range array{
        z01.PrintRune(s)
    }
    z01.PrintRune('\n')
}