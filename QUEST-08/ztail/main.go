package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args[1:]

	intfrombytes := Atoi(arguments[2])
	var bytes []byte
	var tempByte byte = 0
	for i := 0; i < intfrombytes; i++ {
		bytes = append(bytes, tempByte)
	}
	file, err := os.Open(arguments[3])
	if err != nil {
		fmt.Println(err.Error() + "\n")
		os.Exit(0)
	} else {
		_, error := file.Read(bytes)
		if error != nil {
			fmt.Println(error.Error() + "\n")
			os.Exit(0)
		} else {
			fmt.Println(string(bytes))
			defer file.Close()
		}
	}
	os.Exit(0)

}

func Atoi(s string) int {

	runes := []rune(s)
	LenRune := 0
	result := 0
	for i := range runes {
		LenRune = i + 1
	}
	if LenRune == 0 {
		return 0
	}

	tens := 1
	for k := 0; k < LenRune-1; k++ {
		if runes[k] == '+' || runes[k] == '-' {
			continue
		}
		tens *= 10
	}

	for i := range runes {
		if (runes[0] == '-' || runes[0] == '+') && i == 0 {
			continue
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		numb := 0
		for j := '0'; j < runes[i]; j++ {
			numb++
		}
		result += numb * tens
		tens /= 10
	}
	if runes[0] == '-' {
		result *= -1
	}
	return result
}

// tail-c 16 = 16 bytes are displayed                   prosto tail = 10 last lines of a file
