package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	Astr := []rune(str)
	for _, v := range Astr {
		z01.PrintRune(v)
	}
}
