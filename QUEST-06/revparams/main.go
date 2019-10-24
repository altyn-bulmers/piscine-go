package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	arguments := os.Args
	var length int
	for l := range arguments {
		length = l
	}

	for i := length; i > 0; i-- {
		for _, j := range arguments[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}

}
