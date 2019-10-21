package main

import (
	"github.com/01-edu/z01"
)

func main(){
	Raid1d(5,4)
}

func Raid1d(x, y int) {
	if x <= 0 || y <= 0 {
		return
	} else {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if j == 0 && (i == y-1 || i == 0) {
					z01.PrintRune('A')
				} else if j == x-1 && (i == y-1 || i == 0) {
					z01.PrintRune('C')
				} else if i != 0 && j != 0 && i != y-1 && j != x-1 {
					z01.PrintRune(' ')
				} else {
					z01.PrintRune('B')
				}

			}
			z01.PrintRune('\n')
		}
	}
}