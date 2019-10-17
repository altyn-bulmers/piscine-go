package piscine

import (
	"fmt"
)

func PrintStr(str string) {
	Astr := []rune(str)
	count := 0
	for i := range Astr {
		i++
		count++
	}
	fmt.Println(count)

}
