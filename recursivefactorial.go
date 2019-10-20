package main

import "fmt"

func main(){
	fmt.Println(RecursiveFactorial(2))
}

func RecursiveFactorial(nb int) int {
	if nb > 1 {
		return RecursiveFactorial(nb -1) * nb
	} else {
		return 1
	}
}


