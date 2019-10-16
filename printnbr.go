package piscine

import (
	"github.com/01-edu/z01"
)

func check(r int) {

	c := '0'
	if r == 0 {
		z01.PrintRune(c)
		return
	}
	for i := 1; i <= r%10; i++ {
		c++
	}
	for i := -1; i >= r%10; i-- {
		c++
	}
	if r/10 != 0 {
		check(r / 10)
	}
	z01.PrintRune(c)
	return

}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	check(n)
}
