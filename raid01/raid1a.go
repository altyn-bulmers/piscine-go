package piscine

import (
	"github.com/01-edu/z01"
)

func Raid1a(x, y int) {

	if x > 0 && y > 0 {

		if x == 1 && y == 1 {
			z01.PrintRune('o')
			z01.PrintRune('\n')
			return
		} else if x == 1 {
			z01.PrintRune('o')
			for t := 1; t < y-1; t++ { //stolko raz skolko y - 2
				if t <= 0 {
					break
				} else {
					z01.PrintRune('\n')
					z01.PrintRune('|')

				}
				if y == 1 {
					z01.PrintRune('\n')
					z01.PrintRune('o')
					return
				}
			}
			z01.PrintRune('\n')
			z01.PrintRune('o')
			z01.PrintRune('\n')

		} else if y == 1 {
			z01.PrintRune('o')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('-') // 3 dash iz 5
			}
			z01.PrintRune('o')
			z01.PrintRune('\n')
			return
		} else {
			z01.PrintRune('o')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('-') // 3 dash iz 5
			}
			z01.PrintRune('o')
			z01.PrintRune('\n')

			for t := 1; t < y-1; t++ { //stolko raz skolko y - 2
				z01.PrintRune('|')

				for j := 1; j < x-1; j++ {
					z01.PrintRune(' ') // 3 pustoty iz 5
				}
				z01.PrintRune('|')
				z01.PrintRune('\n')

			}

			z01.PrintRune('o')

			for i := 1; i < x-1; i++ {
				z01.PrintRune('-') // 3 dash iz 5
			}
			z01.PrintRune('o')
			z01.PrintRune('\n')

		}
	}
}
