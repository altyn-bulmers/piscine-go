package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args
	args = args[1:]
	len := 0
	for i := range args {
		len = i + 1
	}
	if len == 0 || args[0] == "-h" || args[0] == "--help" {
		fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
	} else {
		var runes []rune
		order := false
		order_prev := false
		insert := false
		for _, arg := range args {
			length := 0
			for j := range arg {
				length = j + 1
			}
			if length > 9 && arg[:9] == "--insert=" {
				runes_temp := []rune(arg[9:])

				for i := 0; i < length-9; i++ {
					runes = append(runes, runes_temp[i])
				}
				insert = true

			} else if length > 3 && arg[:3] == "-i=" {
				runes_temp := []rune(arg[3:])

				for i := 0; i < length-3; i++ {
					runes = append(runes, runes_temp[i])
				}
				insert = true

			} else if (length == 7 && arg[:7] == "--order") || (length == 2 && arg[:2] == "-o") {
				order = true
				order_prev = true
				continue
			} else if order_prev {
				runes_temp := []rune(arg)
				for i := range runes_temp {
					runes = append(runes, runes_temp[i])
				}
			} else {
				runes_print := []rune(arg)
				if order {
					for i := range runes_print {
						runes = append(runes, runes_print[i])
					}
				}
				for i := range runes_print {
					z01.PrintRune(runes_print[i])
				}
			}
			order_prev = false

		}
		if order {
			Order(runes)
		}

		if order || insert {
			for i := range runes {
				z01.PrintRune(runes[i])
			}
		}

		z01.PrintRune(10)
	}

}

func Order(array []rune) {
	len := 0
	for index := range array {
		len = index + 1
	}

	sorted := false
	for !sorted {
		count := 0
		for i := 1; i < len; i++ {
			if array[i] < array[i-1] {
				array[i], array[i-1] = array[i-1], array[i]
				count++
			}
		}
		if count == 0 {
			sorted = true
		}
	}
}
