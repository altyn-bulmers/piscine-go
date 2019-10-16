package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	// count number of digits in n (ex: 123 -> 3 digits
	count := 0
	for c := n; c != 0; count++ {
		c = c / 10
	}

	if count == 0 {
		z01.PrintRune('0')
	} else if n != 0 {
		elem := n
		if n < 0 && n != -9223372036854775808 {
			// convert negatives number to positives & display '-'
			elem = -n
			z01.PrintRune('-')
		} else if n == -9223372036854775808 {
			elem = -n - 1
			z01.PrintRune('-')
		}
		for i := count; i > 0; i-- {
			power := 1
			for d := i - 1; d > 0; d-- {
				power = power * 10
			}
			// treat special limit of 9223372036854775808
			if n == -9223372036854775808 && power == 1 {
				z01.PrintRune(rune(((elem / power) % 10) + 49))
			} else if !(n == -9223372036854775808 && power == 1) {
				z01.PrintRune(rune(((elem / power) % 10) + 48))
			}
		}
	}
}
