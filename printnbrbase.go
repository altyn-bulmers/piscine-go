package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	len := StrLength(base)
	runes := []rune(base)
	valid := true
	if len < 2 {
		valid = false
	} else {
		for i := 0; i < len; i++ {
			if runes[i] == '-' || runes[i] == '+' {
				valid = false
			}
			for j := i + 1; j < len; j++ {
				if runes[i] == runes[j] {
					valid = false
				}
			}
		}
	}
	if !valid {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		if nbr == 0 {
			z01.PrintRune(runes[0])
		} else {
			if nbr < 0 {
				z01.PrintRune('-')
			}
			BaseRecursion(nbr, runes, len)
		}
	}
}

func BaseRecursion(n int, runes []rune, len int) {
	if n/len != 0 {
		BaseRecursion(n/len, runes, len)
	}
	mod := n % len
	if mod < 0 {
		mod = -mod
	}
	z01.PrintRune(runes[mod])
}

func StrLength(str string) int {
	runes := []rune(str)
	var count int = 0
	for i := range runes {
		count++
		i++
	}
	return count
}
