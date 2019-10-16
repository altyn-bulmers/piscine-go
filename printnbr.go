package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	s := 1
	if n < 0 {
		s = -1
		z01.PrintRune('-')
	}

	if n != 0 {
		f := (n / 10) * s
		if f != 0 {
			PrintNbr(f)
		}
		k := ((n % 10) * s) + '0'
		z01.PrintRune(rune(k))

	} else {
		z01.PrintRune('0')
	}
}














