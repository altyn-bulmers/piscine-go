package main

import "github.com/01-edu/z01"

func main(){

	z01.PrintRune('o')

}

func Raid1a(x, y int) {	
	k = []rune(x)
	e = []rune(y)

	lengthk := 0
	lengthe := 0 
	for i :=  range k{
		lengthk = i
	}
	for j := range e{
		lengthe = i
	}


	z01.PrintRune('o')


	for i:= range k - 2 {
		z01.PrintRune('-') // 3 dash iz 5
	}

	for j:= range k -2 {
		z01.PrintRune(' ') // 3 pustoty iz 5
	}

}