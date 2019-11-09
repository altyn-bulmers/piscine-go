package main

import "fmt"

func main() {
	for i := 1; i <=26; i++{
		numb := (14 + i) % 26
		fmt.Println("index: ", i, "number: ", numb)
	}
	
}