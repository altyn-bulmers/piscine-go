package main

import (
	"fmt"
)

const (
	Alphabet = "abcdefghijklmnopqrstuvwxyz"
)


func main() {
        fmt.Printf("%s ", Alphabet)


}

/*func loweralpha() string {
	p := make([]byte, 26)
	for i := range p {
		p[i] = 'a' + byte(i)
	}
	return string(p)
}*/
