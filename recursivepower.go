package main

import "fmt"

func main(){
	fmt.Println(RecursivePower(3,0))
}

func RecursivePower(nb int, power int) int {
	if power < 1 {
		return 1
	} else {

		return RecursivePower(nb, power-1) * nb 


	}

	
}
