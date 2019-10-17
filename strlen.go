package piscine

import (
	"fmt"
)

func StrLen(str string) {
	Astr := []rune(str)
	count := 0
	for i := range Astr {
		i++
		count++
	}
	fmt.Printf("%v", count)

}
