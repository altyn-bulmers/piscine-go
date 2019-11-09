package main

import (
	"fmt"
//	"github.com/01-edu/z01"
)

func main() {
	str := "Altynbek"
	fmt.Println(string(FirstRune(str)))
}

func FirstRune(s string) rune {
	runes := []rune(s)
	return runes[0]
}