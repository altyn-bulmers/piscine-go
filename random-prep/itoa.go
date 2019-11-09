package main

import (
	"github.com/01-edu/z01"
//	"fmt"
)

func main() {
	nbr := -12345 //integer
	input := -nbr 
	nbr = - nbr

	len := 0 
	for input != 0 {
		input /= 10
		len += 1
	}

	array := make([]rune, len + 1) // +1 because of the minus sign
	
	for i := len; i > 0; i--{
		array[i] = rune((nbr % 10) + '0') 
		nbr /= 10	
	}
	array[0] = '-'
	for _,s := range array{
		z01.PrintRune(s)
	}
	z01.PrintRune(10)
//	fmt.Println(array)
}