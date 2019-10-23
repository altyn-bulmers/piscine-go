package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	arguments := os.Args

	for i := range arguments {
		for j := range arguments[i] {
			runes := []rune(arguments[i])
			if i != 0 {
				z01.PrintRune(runes[j])
			}
		}
		z01.PrintRune('\n')
	}
}
