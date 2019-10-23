package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	arguments := os.Args
	name := arguments[0]

	runes := []rune(name)
	for i := range runes {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}
