package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	arguments := os.Args[1:]
	lengthOfArg := 0
	for i := range arguments {
		lengthOfArg = i + 1
	}

	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

func printStr(str string) {
	arrayStr := []rune(str)

	length := 0
	for j := range arrayStr {
		length = j + 1
	}

	for i := 0; i < length; i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}
