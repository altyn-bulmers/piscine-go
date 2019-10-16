package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := i; k <= '9'; k++ {
				for f := j + 1; f <= '9'; f++ {

					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(f)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}

	}
	z01.PrintRune('\n')
}
