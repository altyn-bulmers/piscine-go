package main

import "fmt"

func Fibonacci(index int) int {
	
	if index >= 3{
		return Fibonacci(index-1) + Fibonacci(index-2)

	} else if index == 0 {
		return 0
	} else if index < 0 {
		return -1
	}
	return 1
}


func main(){
	fmt.Println(Fibonacci (5))	
//	1 1 2 3 5 8 13 21 34 55 89 144 233 
}  
