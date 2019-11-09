package main

import (
	"fmt"
	"os"
)

func main() {
	words := []string{"01", "galaxy", "galaxy 01"}
	count := 0
	arguments := os.Args[1:]

	for i := range arguments {
		for _, s := range words {
			if arguments[i] == s {
				count++
			}
		}
	}
	if count >= 1 {
		fmt.Println("Alert!!!")
	}
}
