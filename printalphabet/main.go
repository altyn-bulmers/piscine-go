package main

import (
	"fmt"
)

const (
	Alphabet = "abcdefghijklmnopqrstuvwxyz"
)


func main() {
	
	for i := 97; i <= 122; i++ {
		fmt.Print(string(i))
	}
	fmt.Print("\n");
}

/*func loweralpha() string {
	p := make([]byte, 26)
	for i := range p {
		p[i] = 'a' + byte(i)
	}
	return string(p)
}*/
