package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	argruments := os.Args[1:]
	length := 0
	for i := range argruments {
		length = i + 1
	}

	if length != 0 {
		str := ""
		first := false

		for _, arg := range argruments {
			if first {
				str += " "
			}
			first = true
			str += arg
		}

		runes := []rune(str)
		var pos []int
		var vow []rune

		for i, r := range runes {
			if r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' || r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
				pos = append(pos, i)
				vow = append(vow, runes[i])
			}
		}
		swap (vow)

		for i := range pos {
			runes[pos[i]] = vow [i]
		}

		for _, r := range runes {
			z01.PrintRune(r)
		}

	}
	z01.PrintRune('\n')

}




func swap (vow []rune) {
	len := 0
	for i := range vow {
		len = i + 1
	}
	for i, j := 0, len -1; i < j; i, j = i +1, j -1 {
		temp := vow[i]
		vow[i] = vow[j]
		vow[j] = temp
	} 
}