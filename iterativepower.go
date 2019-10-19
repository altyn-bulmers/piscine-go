package main

import "fmt"

func main(){
	fmt.Println(IterativePower(4,3))
}

func IterativePower(nb int, power int) int {
	result := 1
	
	for i:=0; i < power; i++{
		result *= nb  
	}
		
	return result
}



