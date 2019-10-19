package main

import "fmt"

func main(){
	fmt.Println(IterativeFactorial(5))
}

func IterativeFactorial(nb int) int {
	
	result := 1
	for i:=1; i < nb+1; i++{
		result *= i
	}
	return result
}

