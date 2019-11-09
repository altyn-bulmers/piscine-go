package main

import (
	"github.com/01-edu/z01"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]
	length := 0
	for l := range arguments {
		length = l + 1
	}

	if length == 0 {
		input, err := ioutil.ReadAll(os.Stdin)
		for _, j := range string(input) {
			z01.PrintRune(j)
		}
		if err != nil {
			for _, e := range err.Error() {
				z01.PrintRune(e)
			}
			z01.PrintRune('\n')
		}
		return
	}

	first := true

	for _, arg := range arguments {
		file, err := os.Open(arg)

		if err != nil {
			for _, e := range err.Error() {
				z01.PrintRune(e)
			}
			z01.PrintRune('\n')
		}
		return

		f, err := ioutil.ReadAll(file)

		if !first {
			z01.PrintRune('\n')
		}
		first = false
		for _, text := range string(f) {
			z01.PrintRune(text)
		}
		if err != nil {
			for _, e := range err.Error() {
				z01.PrintRune(e)
			}
			z01.PrintRune('\n')
		}

		file.Close()
	}

}
