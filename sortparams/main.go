package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args

	for i := range arguments {
		runes := []rune(arguments[i])
		j := 0
		if i != 0 && (runes[j] >= '0' && runes[j] <= '9') {
			for _, j := range arguments[i] {
				z01.PrintRune(j)
				z01.PrintRune('\n')
			}
		}
	}
	for i := range arguments {
		runes := []rune(arguments[i])
		j := 0
		if i != 0 && (runes[j] >= 'A' && runes[j] <= 'Z') {
			for _, j := range arguments[i] {
				z01.PrintRune(j)
				z01.PrintRune('\n')
			}
		}
	}
	for i := range arguments {
		runes := []rune(arguments[i])
		j := 0
		if i != 0 && (runes[j] >= 'a' && runes[j] <= 'z') {
			for _, j := range arguments[i] {
				z01.PrintRune(j)
				z01.PrintRune('\n')
			}
		}
	}

}
