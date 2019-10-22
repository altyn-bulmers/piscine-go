package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	var s []int

	if n == 0 {
		s = append(s, 0)
	}
	for i := 1; n > 0; i++ {
		s = append(s, n%10)
		n /= 10
	}
	Sort(s)
	for i := range s {
		z01.PrintRune(rune('0' + s[i]))
	}
}

func Sort(table []int) {

	var length int = 0
	for i := range table {
		i++
		length++
	}

	i := 1
	for i < length {
		if table[i-1] > table[i] {
			tmp := table[i-1]
			table[i-1] = table[i]
			table[i] = tmp
			i = 1
		} else {
			i++

		}
	}
}
