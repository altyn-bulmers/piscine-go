package main

import "fmt"

func main(){
	fmt.Println(RecursiveFactorial(8))
}

func RecursiveFactorial(nb int) int {
	if nb >=3 {
		return RecursiveFactorial(nb -1) + RecursiveFactorial(nb -2)
	} else {
		return 1
	}
}


