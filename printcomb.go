package piscine

import "github.com/01-edu/z01"

func PrintComb() {	
	next := false
	for i :=0; i <= 9; i++ {
        for j := i+1; j <= 9; j++{
            for k := j+1; k <= 9; k++ {
				
				if next {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				next = true
				z01.PrintRune(rune(i+48))
				z01.PrintRune(rune(j+48))
				z01.PrintRune(rune(k+48))
			
            }
        }
    }
	z01.PrintRune('\n')
	return 


}
