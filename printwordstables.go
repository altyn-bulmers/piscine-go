package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	length := 0

	for i := range table {
		length = i + 1
	}

	for _, v := range table {
		Printrune(v)
	}
}

func Printrune(str string) {
	strRune := []rune(str)

	for _, v := range strRune {
		z01.Printrune(v)
	}
	z01.Printrune('\n')

}
