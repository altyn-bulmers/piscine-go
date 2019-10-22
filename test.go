package main

import (
//	"github.com/01-edu/z01"
	"fmt"
)	

func main(){
	s := "Ola!"
	toFind := "hOl"
	runes := []rune(s)
	sub := []rune(toFind)

	length := 0
	sublength := 0

	for i := range s{
		length = i + 1
	}

	for i := range toFind{
		sublength = i + 1
	}
	if length < sublength {
		fmt.Println("-1")
	}
	if length == 0 {
		fmt.Println("0")
	}
	for i := range runes {
		if runes[i] == sub[0] {
			fmt.Println(i)
		}
		
	}

//	fmt.Println("lenght:", length)

	
//	z01.PrintRune(rune(runes[length-1]))

//	z01.PrintRune('\n')
//	Atoi(runes)
	
}
