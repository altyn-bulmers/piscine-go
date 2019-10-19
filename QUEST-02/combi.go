package main

import (
	
	"github.com/01-edu/z01"

)
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

func PrintComb123() {	
	next := false
	
	for i :='0'; i <= '9'; i++ {
        for j := i+1; j <= '9'; j++{
            for k := j+1; k <= '9'; k++ {
				if next{
				z01.PrintRune(',')
				z01.PrintRune(' ')
				}
				
				next = true
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
			}
				

					
		}
				
			
    }
	z01.PrintRune('\n')
	return
	    
}
	 

func PrintComb333() {

	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			f := j + 1
			for k := i; k <= '9'; k++ {
				for ; f <= '9'; f++ {
					z01.PrintRune((i))
					z01.PrintRune((j))
					z01.PrintRune(' ')
					z01.PrintRune((k))
					z01.PrintRune((f))
					if i < '9' || j < '8' || k < '9' || f < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
				f = '0'
			}
		}
	}
	z01.PrintRune('\n')

}

func main () {
	PrintComb333()

}

