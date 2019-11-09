package main

import "github.com/01-edu/z01"

func main() {
	base := "0123456789ABCDEF"
	nbr := 125
	len := strLen(base)
	runes := []rune(base)
	valid := true
	if len < 2 {
		valid = false
	} else {
		for i := 0; i < len; i++ {
			if runes[i] == '-' || runes[i] == '+' {
				valid = false
			}
			for j := i+ 1; j < len; j++{
				if runes[i] == runes[j]{
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
			z01.PrintRune('0')
		} else {
			if nbr <0 {
				z01.PrintRune('-')
			}
			BaseRecursion(nbr, runes, len)
		}
	}
	z01.PrintRune('\n')

}

func BaseRecursion(n int, runes []rune, len int) {
	if n/len != 0 {
		BaseRecursion(n/len, runes, len)
	}
	mod := 0
	mod = n % len 
	if mod < 0 {
		mod = -mod
	}
	z01.PrintRune(runes[mod])
}


func strLen(s string) int {
	count := 0
	for i := range s {
		count = i + 1
	}
	return count
}