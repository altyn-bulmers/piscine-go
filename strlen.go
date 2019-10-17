package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	Astr := []rune(str)
	for i, v := range Astr {
		z01.PrintRune(i)
	}
}
