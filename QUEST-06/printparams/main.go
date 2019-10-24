package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	arguments := os.Args

	for i := range arguments {
		if i != 0 {
			for j := range arguments[i] {
				runes := []rune(arguments[i])
				{
					z01.PrintRune(runes[j])
				}
			}
			z01.PrintRune('\n')
		}
	}
}
